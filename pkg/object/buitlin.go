package object

// Builtin represents builtin function
type Builtin struct {
	Function BuiltinFunction
}

func (b *Builtin) Type() ObjectType {
	return BuiltinObject
}

func (b *Builtin) Inspect() string {
	return "Lail Builtin Function"
}
