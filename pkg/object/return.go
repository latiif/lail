package object

type Return struct {
	Value Object
}

func (r *Return) Inspect() string {
	return r.Value.Inspect()
}

func (r *Return) Type() ObjectType {
	return ReturnObject
}
