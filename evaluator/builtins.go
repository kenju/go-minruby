package evaluator

import (
	"fmt"
	"github.com/kenju/go-minruby/object"
)

var builtins = map[string]*object.Builtin{
	"p": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			fmt.Printf("+%v\n", args)

			return &object.Null{}
		},
	},
}