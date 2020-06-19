package ast

import (
	"bytes"

	"github.com/latiif/lail/pkg/token"
)

// LetStatement defines the 'let <id> = <expr>';
type LetStatement struct {
	Token token.Token // the token.Let token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

// TokenLiteral implements the Node interface
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
