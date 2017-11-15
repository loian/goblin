package ast

import (
	"goblin/token"
)

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Types      []uint16
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) GetTypes() []uint16   { return es.Types }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
