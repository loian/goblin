package ast

import (
	"bytes"

	"goblin/token"
)

type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Types    []uint16
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) GetTypes() []uint16   { return pe.Types }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())

	out.WriteString(")")

	return out.String()
}
