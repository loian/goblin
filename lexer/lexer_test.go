package lexer

import (
	"testing"

	"goblin/token"
)

func TestNextToken(t *testing.T) {
	input := `
		var word string;
		let number = 34.1;

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
