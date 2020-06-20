package interpretor

import "github.com/latiif/lail/pkg/object"

var builtins = map[string]*object.Builtin{
	"len": {
		Function: func(args ...object.Object) object.Object {
			return Null
		},
	},
}
