package ast

import (
	"goblin/token"
)



type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
	GetTypes() []uint16
}

type Expression interface {
	Node
	GetTypes() []uint16
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}


type Identifier struct {
	Token token.Token // the token.IDENT token
	Types []uint16
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i* Identifier) GetTypes() []uint16 {return i.Types}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }


