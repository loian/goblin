package lexer

import (
	"goblin/token"

	"errors"
)

//Lexer
type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
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

func (l *Lexer) peekRune() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readStringToken() (string, error) {
	position := l.position
	var lastRune rune
	//is the first rune?
	afterFirst := false

	//string tokens (like identifiers and keywords) must start with letters
	for isLetter(l.ch) || afterFirst && isDigit(l.ch) || afterFirst && l.ch == '.' {
		if !afterFirst {
			afterFirst = true
		}
		lastRune = l.ch
		l.readRune()
	}

	//strings token can't end with .

	var err error
	if lastRune == '.' {
		err = errors.New("string token can't terminate with a dot.")
	}

	return string(l.input[position:l.position]), err
}

func (l *Lexer) readNumber() (token.TokenType, string) {
	isDecimal := false
	position := l.position

	for l.ch == '.' || isDigit(l.ch) {
		if l.ch == '.' {
			isDecimal = true
		}
		l.readRune()
	}

	numberType := token.TokenType(token.INT)
	if isDecimal == true {
		numberType = token.TokenType(token.DECIMAL)
	}
	return numberType, string(l.input[position:l.position])
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekRune() == '=' {
			ch := l.ch
			l.readRune()
			literal := string(ch) + string(l.ch)
			tok = token.Token{token.EQUAL, literal}
		} else {
			tok = token.Token{token.ASSIGN, string(l.ch)}
		}

	case '!':
		if l.peekRune() == '=' {
			ch := l.ch
			l.readRune()
			literal := string(ch) + string(l.ch)
			tok = token.Token{token.NOTEQUAL, literal}
		} else {
			tok = token.Token{token.NOT, string(l.ch)}
		}

	case ';':
		tok = token.Token{token.SEMICOLON, string(l.ch)}
	case '(':
		tok = token.Token{token.LPAREN, string(l.ch)}
	case ')':
		tok = token.Token{token.RPAREN, string(l.ch)}
	case ',':
		tok = token.Token{token.COMMA, string(l.ch)}
	case '+':
		tok = token.Token{token.PLUS, string(l.ch)}
	case '-':
		tok = token.Token{token.MINUS, string(l.ch)}
	case '*':
		tok = token.Token{token.MULTIPLY, string(l.ch)}
	case '/':
		tok = token.Token{token.DIVIDE, string(l.ch)}
	case '>':
		if l.peekRune() == '=' {
			ch := l.ch
			l.readRune()
			literal := string(ch) + string(l.ch)
			tok = token.Token{token.GREATEREQUAL, literal}
		} else {
			tok = token.Token{token.GREATER, string(l.ch)}
		}
	case '<':
		if l.peekRune() == '=' {
			ch := l.ch
			l.readRune()
			literal := string(ch) + string(l.ch)
			tok = token.Token{token.LESSEREQUAL, literal}
		} else {
			tok = token.Token{token.LESSER, string(l.ch)}
		}
	case '{':
		tok = token.Token{token.LBRACE, string(l.ch)}
	case '}':
		tok = token.Token{token.RBRACE, string(l.ch)}
	case 0:
		tok = token.Token{token.EOF, ""}

	default:

		// if the token starts with a letter
		if isLetter(l.ch) {

			//Build the token literal
			literal, err := l.readStringToken()
			if err != nil {
				return token.Token{token.ILLEGAL, literal}
			}

			//now we have to choose the token tyoe
			//it can be a type name
			tokenType, err := token.LookupType(literal)

			//or as a last option an identifier
			if err != nil {
				tokenType = token.IDENT
			}

			//finally return the token
			return token.Token{tokenType, literal}

		} else if isDigit(l.ch) {
			// it could be a number

			t, l := l.readNumber()
			tok.Type = t
			tok.Literal = l
			return tok
		}

		tok = token.Token{token.ILLEGAL, string(l.ch)}

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
