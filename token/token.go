package token

const (
	ILLEGAL = iota
	EOF     = 1
)

const (
	// Identifiers + literals
	IDENT = iota + 100
)

const(
	//BASIC TYPES
	TYPENAME = iota + 200
	INT
	FLOAT
	DECIMAL
	STRING
)

const(
	// ASSIGNMENT
	ASSIGN         = iota + 300
	ASSIGNPLUS
	ASSINGMINUS
	ASSIGNMULTIPLY
	ASSIGNDIVIDE
)

const (
	//OPERATORS
	PLUS         = iota + 400
	MINUS
	MULTIPLY
	DIVIDE
	MODULE
	GREATER
	GREATEREQUAL
	LESSER
	LESSEREQUAL
)

const (
	// Delimiters
	COMMA     = iota + 500
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET
)

const (
	// Keywords
	FALSE    = iota + 600
	FUNCTION
	LET
	TRUE
	VAR

)

//Defines the type of the tokens
type TokenType uint16

//Token a structure to holds token informations
type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, rn rune) Token {
	return Token{Type: tokenType, Literal: string(rn)}
}
