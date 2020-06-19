package object

import (
	"bytes"
	"strings"

	"github.com/latiif/lail/pkg/ast"
)

// Function represents a Function object
type Function struct {
	Params []*ast.Identifier
	Body   *ast.BlockStatement
	Env    *Env
}

func (f *Function) Type() ObjectType {
	return FunctionObject
}

func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Params {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}
