package token

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"var": VAR,
}

var typenames = map[string]TokenType {
	"int": TYPENAME,
	"string": TYPENAME,
	"float": TYPENAME,
	"decimal": TYPENAME,
}

func LookupType(t string) (TokenType, bool) {
	if tok, ok := typenames[t]; ok {
		return tok, true
	}
	return ILLEGAL, false
}



func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

