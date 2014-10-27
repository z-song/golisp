package ext

import (
	"../compiler"
	//"fmt"
	//"strings"
)

// (fill [] 3 "hello")  -> ["hello" "hello" "hello"]
func Fill(args []compiler.Node, Env compiler.Environment) (ret interface{}) {
	if len(args) != 3 {
		panic("[function fill] function need 3 arguments")
	}

	if args[0].Type != compiler.Tarray {
		panic("[function fill] argument 1 should be array type")
	}

	var res []interface{}
	for i := 0; i < args[1].ToInt(); i++ {
		res = append(res, args[2].Value())
	}

	return compiler.NewNode(res)
}

func Member(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

}
