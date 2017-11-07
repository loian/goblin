package lexer

import (
	"testing"

	"goblin/token"
)



func TestNextToken(t *testing.T) {
	input := `
		var word string;
		let number = 34.1;
		let number2 float = 14;
		let identIllegal. = 14;
		let myFunc = fn(param string) string, int {};
		a == 3
	`


	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "word"},
		{token.TYPENAME, "string"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "number"},
		{token.ASSIGN, "="},
		{token.FLOAT, "34.1"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "number2"},
		{token.TYPENAME, "float"},
		{token.ASSIGN, "="},
		{token.INT, "14"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.ILLEGAL, "identIllegal."},
		{token.ASSIGN, "="},
		{token.INT, "14"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "myFunc"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "param"},
		{token.TYPENAME, "string"},
		{token.RPAREN, ")"},
		{token.TYPENAME, "string"},
		{token.COMMA, ","},
		{token.TYPENAME, "int"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "a"},
		{token.EQUAL, "=="},
		{token.INT, "3"},

	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%s, got=%s",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%c, got=%c",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
