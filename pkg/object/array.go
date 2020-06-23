package object

import (
	"bytes"
	"strings"
)

type Array struct {
	Value []Object
}

func (a *Array) Type() ObjectType {
	return ArrayObject
}

func (a *Array) Inspect() string {
	var out bytes.Buffer
	elements := make([]string, len(a.Value))
	out.WriteString("[")

	for i, v := range a.Value {
		elements[i] = v.Inspect()
	}
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
