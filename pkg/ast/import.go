package ast

import (
	"github.com/latiif/lail/pkg/token"
)

// ImportStatement returns a program node
type ImportStatement struct {
	Token   token.Token
	Program *Program
}

func (is *ImportStatement) statementNode() {}

// TokenLiteral implements the Node interface
func (is *ImportStatement) TokenLiteral() string {
	return is.Token.Literal
}

func (is *ImportStatement) String() string {
	return is.Program.String()
}
