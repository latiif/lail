package ast

import (
	"bytes"
	"strings"

	"github.com/latiif/lail/pkg/token"
)

// Array represents an array of elements regardless of type
type Array struct {
	Token    token.Token
	Elements []Expression
}

func (a *Array) expressionNode() {}

// TokenLiteral implements the Node interface
func (a *Array) TokenLiteral() string {
	return a.Token.Literal
}

func (a *Array) String() string {
	var out bytes.Buffer
	elements := make([]string, len(a.Elements))
	out.WriteString("[")

	for i, v := range a.Elements {
		elements[i] = v.String()
	}
	out.WriteString(strings.Join(elements, ","))
	out.WriteString("]")

	return out.String()
}
