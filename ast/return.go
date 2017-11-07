package ast

import (
	"goblin/token"
)

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) GetTypes() []uint16 {return rs.ReturnValue.GetTypes()}
