package ast

import "github.com/latiif/lail/pkg/token"

// Identifier represents a valid identifier
type Identifier struct {
	Token token.Token // The token.Ident token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral implements the Node interface
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
