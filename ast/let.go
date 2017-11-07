package ast

import (
	"goblin/token"
)

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Types  []uint16
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) GetTypes() []uint16 {return ls.Types}