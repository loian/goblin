package ast

import (
	"goblin/tables"
	"goblin/token"
)

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BooleanLiteral) String() string       { return bl.Token.Literal }
func (bl *BooleanLiteral) GetTypes() []uint16   { return []uint16{tables.BOOL} }
