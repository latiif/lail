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
			array, ok := args[0].(*object.Array)
			if !ok {
				return newIllegalStateException(fmt.Sprintf("head: %s is not an array literal.", args[0].Inspect()))
			}
			// head of [] is Null
			if len(array.Value) == 0 {
				return Null
			}
			return array.Value[0]
		},
	},
	"tail": {
		Function: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newIllegalStateException(fmt.Sprintf("tail takes 1 argument; %d were provided.", len(args)))
			}
			array, ok := args[0].(*object.Array)
			if !ok {
				return newIllegalStateException(fmt.Sprintf("tail: %s is not an array literal.", args[0].Inspect()))
			}
			// return empty array
			// tail of [] is []
			if len(array.Value) == 0 {
				return &object.Array{}
			}
			return &object.Array{
				Value: array.Value[1:],
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
