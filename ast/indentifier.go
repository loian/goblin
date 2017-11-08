package ast

import (
	"goblin/token"
)

type Identifier struct {
	Token token.Token // the token.IDENT token
	Types []uint16
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) GetTypes() []uint16 {return i.Types}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {return i.Value}

