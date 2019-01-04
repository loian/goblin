package ast

type Node interface {
	TokenLiteral() string
	GetTypes() []uint16
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}
