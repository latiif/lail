package ast

import "bytes"

// Node is the smallest building block
type Node interface {
	// TokenLiteral returns the literal value of it's associated token
	TokenLiteral() string
	String() string
}

// Statement nodes
type Statement interface {
	Node
	statementNode()
}

// Expression nodes
type Expression interface {
	Node
	expressionNode()
}

// Program represents a program i.e. a slice of statements
type Program struct {
	Statements []Statement
}

// TokenLiteral recursively prints literals of the program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
