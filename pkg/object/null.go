package object

type Null struct{}

func (n *Null) Type() ObjectType { return NullObject }
func (n *Null) Inspect() string {
	return "null"
}
