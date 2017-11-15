package parser

import (
	"goblin/ast"
	"goblin/tables"
	"goblin/token"
)

//Parse a let
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	//if the next token is an ident or a basic type name, assign the type to the statement
	if p.peekToken.Type == token.IDENT || p.peekToken.Type == token.TYPENAME {
		t, err := tables.LookupTypeCode(p.peekToken.Literal)

		if err != nil {
			//store the type literal to late-resolving it
			tables.LateResolveType(p.peekToken.Literal)
		}

		//The given identifier points to a type, so we can declare the let as accepting only that specific type
		stmt.Types = append(stmt.Types, t)
		//and then move to the next token
		p.nextToken()
	}

	//the next token has to be the =
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt

}
