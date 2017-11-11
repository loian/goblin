package token

import "errors"

var keywords = map[string]TokenType{
	"else":   ELSE,
	"false":  FALSE,
	"fn":     FUNCTION,
	"if":     IF,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
	"var":    VAR,
}

var typenames = map[string]TokenType{
	"bool":    TYPENAME,
	"int":     TYPENAME,
	"string":  TYPENAME,
	"float":   TYPENAME,
	"decimal": TYPENAME,
}

var tokenNames = map[TokenType]string{
	ILLEGAL:        "ILLEGAL",
	EOF:            "EOF",
	IDENT:          "IDENT",
	TYPENAME:       "TYPENAME",
	BOOL:           "BOOL",
	INT:            "INT",
	DECIMAL:        "DECIMAL",
	STRING:         "STRING",
	NULL:           "NULL",
	ALL:            "ALL",

	ASSIGN:         "ASSIGN",
	ASSIGNPLUS:     "ASSIGNPLUS",
	ASSINGMINUS:    "ASSINGMINUS",
	ASSIGNMULTIPLY: "ASSIGNMULTIPLY",
	ASSIGNDIVIDE:   "ASSIGNDIVIDE",
	PLUS:           "PLUS",
	MINUS:          "MINUS",
	MULTIPLY:       "MULTIPLY",
	DIVIDE:         "DIVIDE",
	MODULE:         "MODULE",
	GREATER:        "GREATER",
	GREATEREQUAL:   "GREATEREQUAL",
	LESSER:         "LESSER",
	LESSEREQUAL:    "LESSEREQUAL",
	NOT:            "NOT",
	EQUAL:          "EQUAL",
	NOTEQUAL:       "NOTEQUAL",
	COMMA:          "COMMA",
	SEMICOLON:      "SEMICOLON",
	LPAREN:         "LPAREN",
	RPAREN:         "RPAREN",
	LBRACE:         "LBRACE",
	RBRACE:         "RBRACE",
	LBRACKET:       "LBRACKET",
	RBRACKET:       "RBRACKET",
	ELSE:           "ELSE",
	FALSE:          "FALSE",
	FUNCTION:       "FUNCTION",
	IF:             "IF",
	LET:            "LET",
	RETURN:         "RETURN",
	TRUE:           "TRUE",
	VAR:            "VAR",
	CONST:          "CONST",
}

func LookupName(tokenType TokenType) string {
	if name, ok := tokenNames[tokenType]; ok {
		return name
	}
	return "unknown"
}

func LookupType(t string) (TokenType, error) {

	if tokenType, ok := keywords[t]; ok {
		return tokenType, nil
	}

	if tokenType, ok := typenames[t]; ok {
		return tokenType, nil
	}

	return 0, errors.New("token lookup produces no results")
}
