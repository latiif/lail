package interpretor

import (
	"bytes"
	"fmt"

	"github.com/latiif/lail/pkg/object"
)

var builtins = map[string]*object.Builtin{
	"out": {
		Function: func(args ...object.Object) object.Object {
			var out bytes.Buffer

			for _, arg := range args {
				out.WriteString(arg.Inspect())
			}

			fmt.Println(out.String())
			return &object.String{
				Value: out.String(),
			}
		},
	},
	"head": {
		Function: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newIllegalStateException(fmt.Sprintf("head takes 1 argument; %d were provided.", len(args)))
			}
			switch args[0].(type) {
			case *object.Array:
				argument := args[0].(*object.Array)
				// head of [] is Null
				if len(argument.Value) == 0 {
					return Null
				}
				return argument.Value[0]
			case *object.String:
				argument := args[0].(*object.String)
				// head of "" is Null
				if len(argument.Value) == 0 {
					return Null
				}
				return &object.String{Value: fmt.Sprintf("%c", argument.Value[0])}
			default:
				return newIllegalStateException(fmt.Sprintf("head: %s is not an array literal.", args[0].Inspect()))
			}
		},
	},
	"tail": {
		Function: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newIllegalStateException(fmt.Sprintf("tail takes 1 argument; %d were provided.", len(args)))
			}
			switch args[0].(type) {
			case *object.Array:
				// return empty array
				// tail of [] is []
				array := args[0].(*object.Array)
				if len(array.Value) == 0 {
					return &object.Array{}
				}
				return &object.Array{
					Value: array.Value[1:],
				}
			case *object.String:
				str := args[0].(*object.String)
				if len(str.Value) == 0 {
					return &object.String{}
				}
				return &object.String{
					Value: str.Value[1:],
				}
			default:
				return newIllegalStateException(fmt.Sprintf("tail: %s is not an array literal.", args[0].Inspect()))
			}
		},
	},
	"typeof": {
		Function: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newIllegalStateException(fmt.Sprintf("typeof takes 1 argument; %d were provided.", len(args)))
			}
			return &object.String{
				Value: string(args[0].Type()),
			}
		},
	},
}
