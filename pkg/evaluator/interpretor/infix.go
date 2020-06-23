package interpretor

import (
	"fmt"
	"strings"

	"github.com/latiif/lail/pkg/object"
)

func evalAsInteger(operand object.Object) int64 {
	switch operand.Type() {
	case object.IntegerObject:
		return operand.(*object.Integer).Value
	case object.BooleanObject:
		if operand.(*object.Boolean).Value {
			return 1
		}
		return 0
	default:
		return 0
	}
}

// evalInfixMinus evaluates - operator in all its forms
// it returns 0 if type mismatch detected
func evalInfixMinus(lhs, rhs object.Object) object.Object {
	// <int> - <int> = substraction
	if rhs.Type() == object.IntegerObject && lhs.Type() == object.IntegerObject {
		return &object.Integer{
			Value: lhs.(*object.Integer).Value - rhs.(*object.Integer).Value,
		}
	}

	// <str> - <str1> = replace first instance of str1 in str
	if rhs.Type() == object.StringObject && lhs.Type() == object.StringObject {
		return &object.String{
			Value: strings.Replace(lhs.Inspect(), rhs.Inspect(), "", 0),
		}
	}

	return &object.Integer{Value: evalAsInteger(lhs) - evalAsInteger(rhs)}

}

// evalInfixPlus evaluates + operator in all its forms
// 1 + 1 => 2
// true + true + false => 2
// "str" + 1 => "str1"
// [2,4,"foo"] + ["bar", true] = [2,4,"foo","bar",true]
func evalInfixPlus(lhs, rhs object.Object) object.Object {

	// <int> + <int> = addition
	if rhs.Type() == object.IntegerObject && lhs.Type() == object.IntegerObject {
		return &object.Integer{
			Value: lhs.(*object.Integer).Value + rhs.(*object.Integer).Value,
		}
	}

	// <arr> + <arr> = [<arr>, <arr>]
	if rhs.Type() == object.ArrayObject && lhs.Type() == object.ArrayObject {
		return &object.Array{
			Value: append(lhs.(*object.Array).Value, rhs.(*object.Array).Value...),
		}
	}

	// <str> + <int> | <bool> | <null> = concat
	if lhs.Type() == object.StringObject {
		return &object.String{
			Value: fmt.Sprintf("%s%s", lhs.(*object.String).Value, rhs.Inspect()),
		}
	}

	// <int> | <bool> | <null> + <str>  = concat
	if rhs.Type() == object.StringObject {
		return &object.String{
			Value: fmt.Sprintf("%s%s", lhs.Inspect(), rhs.(*object.String).Value),
		}
	}
	return &object.Integer{Value: evalAsInteger(lhs) + evalAsInteger(rhs)}

}

// evalInfixEquality checks for equality regardless of type
// 1 == "1" => true
func evalInfixEquality(lhs, rhs object.Object) object.Object {
	return &object.Boolean{
		Value: lhs.Inspect() == rhs.Inspect(),
	}
}

// evalInfixInequality checks for inequality irragrdless of type
func evalInfixInequality(lhs, rhs object.Object) object.Object {
	return &object.Boolean{
		Value: lhs.Inspect() != rhs.Inspect(),
	}
}
