package lexer

import (
	"goblin/token"

	"fmt"
)

//Lexer
type Lexer struct {
	input []rune
	position int
	readPosition int
	ch rune
}

//New returns a new Lexer pointer
func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readRune()
	return l
}

func (l *Lexer) readRune() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readRune()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readNumber() (token.TokenType, string) {
	isFloat := false
	position := l.position

	for l.ch == '.' || isDigit(l.ch)    {
		if l.ch == '.' {
			isFloat = true
			fmt.Println("FLoat found")
		}
		l.readRune()
	}

	numberType := token.TokenType(token.INT);
	if isFloat == true {
		numberType = token.TokenType(token.FLOAT)
	}
	return numberType, string(l.input[position:l.position])
}


func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace();

	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:

		// if the token starts with a letter
		if isLetter(l.ch) {
			var isType = false;
			tok.Literal = l.readIdentifier()

			//it can be a type name
			tok.Type, isType = token.LookupType(tok.Literal)

			//or as a last option an identifier
			if !isType {
				tok.Type = token.LookupIdent(tok.Literal)
			}
			return tok
		} else if isDigit(l.ch) {
			// it could be a number

			t,l := l.readNumber()
			tok.Type = t
			tok.Literal = l
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readRune()
	return tok
}

func isLetter(rn rune) bool {
	return 'a' <= rn && rn <= 'z' || 'A' <= rn && rn <= 'Z' || rn == '_'
}

func isDigit(rn rune) bool {
	return '0' <= rn && rn <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readRune()
	}
}

