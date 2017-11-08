package token

const (
	ILLEGAL = iota
	EOF     = 1
)

const (
	// Identifiers + reserved
	IDENT = iota + 100
	TYPENAME
)

const(
	//BASIC TYPES
	BOOL  = iota + 200
	INT
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
	NOT
	EQUAL
	NOTEQUAL
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
	CONST     = iota + 600
	ELSE
	FALSE
	FUNCTION
	IF
	LET
	RETURN
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

