package object

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	IntegerObject  = "Integer"
	BooleanObject  = "Boolean"
	NullObject     = "Null"
	ReturnObject   = "ReturnObject"
	FunctionObject = "FunctionObject"
	StringObject   = "StringObject"
	ErrorObject    = "ErrorObject"
)
