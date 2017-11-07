package parser

import (
	"goblin/token"
	"goblin/ast"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: We're skipping the expressions until we
	// encounter a semicolon

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt

}