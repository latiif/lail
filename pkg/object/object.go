package object

type ObjectType string

type BuiltinFunction func(args ...Object) Object

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	IntegerObject  = "Integer"
	BooleanObject  = "Boolean"
	ArrayObject    = "Array"
	NullObject     = "Null"
	ReturnObject   = "ReturnObject"
	FunctionObject = "Function"
	StringObject   = "String"
	ErrorObject    = "Error"
	BuiltinObject  = "BuiltinObject"
)
