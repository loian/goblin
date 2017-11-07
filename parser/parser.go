package parser

import (
	"goblin/lexer"
	"goblin/token"
	"goblin/ast"
	"fmt"
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		Errors: []string{},
		}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

//Parser object
type Parser struct {
	l *lexer.Lexer
	Errors []string

	curToken  token.Token
	peekToken token.Token
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
		return nil
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

func (p *Parser) typeError(ident string) {
	msg := fmt.Sprintf("unexpected identifier %s", ident)
	p.Errors = append(p.Errors, msg)
}