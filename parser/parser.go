package parser

import (
	"goblin/lexer"
	"goblin/token"
	"goblin/ast"
	"fmt"
	"goblin/tables"
	"strconv"
	"github.com/shopspring/decimal"
)

//define operators precedence
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	SUFFIX		// X
	CALL        // myFunction(X)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
	postfixParseFn func() ast.Expression
)

//Constructor function
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		Errors: []string{},
	}

	//Register parser prefix functions
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.DECIMAL, p.parseDecimalLiteral)
	p.registerPrefix(token.NOT, p.parseNotPrefixExpression)
	p.registerPrefix(token.MINUS, p.parseMinusPrefixExpression)

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

//Parser object
type Parser struct {
	l      *lexer.Lexer
	Errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
	postfixParseFns map[token.TokenType]postfixParseFn
}

//Get the next lexer token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

//Parse a goblin program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	lit, e := tables.LateTypeResolvingCheck();
	if e != nil {
		p.typeError(lit)
	}
	tables.LateTypeMapCleanUp()

	return program
}

//Parse the current statement
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

//Resolve the token type of the current token
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

//Resolve the token type of the next token
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

//Assert function, it record an error if the expected type is not the wanted one
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

//Record an error
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		token.LookupName(t), token.LookupName(p.peekToken.Type))
	p.Errors = append(p.Errors, msg)
}

//Record a type error
func (p *Parser) typeError(ident string) {
	msg := fmt.Sprintf("unexpected identifier %s", ident)
	p.Errors = append(p.Errors, msg)
}

//Record a type error
func (p *Parser) dataTypeError(t uint16) {
	tn,_:=tables.LookupTypeName(t)
	msg := fmt.Sprintf("unexpected type %s", tn)
	p.Errors = append(p.Errors, msg)
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) registerPostfix(tokenType token.TokenType, fn postfixParseFn) {
	p.postfixParseFns[tokenType] = fn
}

//Parse an expression
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)

		return nil
	}

	leftExp := prefix()

	return leftExp
}


func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.Errors = append(p.Errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parseDecimalLiteral() ast.Expression {
	lit := &ast.DecimalLiteral{Token: p.curToken}

	value, err := decimal.NewFromString(p.curToken.Literal)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as decimal", p.curToken.Literal)
		p.Errors = append(p.Errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}



func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.Errors = append(p.Errors, msg)
}

//Parse a ! prefix expression
func (p *Parser) parseNotPrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		//a not expression can be applied to these types
		Types: []uint16{tables.BOOL, tables.INT, tables.DECIMAL},
	}
	//move to the next token, after the prefix has been parsed
	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	//Compares the type of the right expression (the one following the current parsed token)
	//with the accepted type of the current expression (a !)
	if !p.checkTypeCompatibility(expression.GetTypes(), expression.Right.GetTypes()) {
		p.dataTypeError(expression.Right.GetTypes()[0])
	}

	return expression
}

//Parse a - prefix expression
func (p *Parser) parseMinusPrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		//a not expression can be applied to these types
		Types: []uint16{tables.INT, tables.DECIMAL},
	}
	//move to the next token, after the prefix has been parsed
	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	//Compares the type of the right expression (the one following the current parsed token)
	//with the accepted type of the current expression (a !)
	if !p.checkTypeCompatibility(expression.GetTypes(), expression.Right.GetTypes()) {
		p.dataTypeError(expression.Right.GetTypes()[0])
	}

	return expression
}

func (p *Parser) checkTypeCompatibility(receiver []uint16, received []uint16) bool{
	for x := range receiver {
		for y:= range received {
			if x==y {
				return true
			}
		}
	}
	return false
}