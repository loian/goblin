package ast

import (
	"goblin/tables"
	"goblin/token"

	"github.com/shopspring/decimal"

)

type DecimalLiteral struct {
	Token token.Token
	Value decimal.Decimal
}

func (il *DecimalLiteral) expressionNode()      {}
func (il *DecimalLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *DecimalLiteral) String() string       { return il.Token.Literal }
func (il *DecimalLiteral) GetTypes() []uint16   { return []uint16{tables.DECIMAL}}

