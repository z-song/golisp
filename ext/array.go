package ext

import (
	"github.com/z-song/golisp/compiler"
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

func Range(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	var arr []interface{}

	//(range 3) -> [0 1 2]
	if len(args) == 1 {
		for i := 0; i < args[0].Value().(int); i++ {
			arr = append(arr, i)
		}
	}

	//(range 3 6) -> [3 4 5]
	if len(args) == 2 {
		for i := args[0].Value().(int); i < args[1].Value().(int); i++ {
			arr = append(arr, i)
		}
	}

	//(range 3 10 1) -> [3 4.5 6 7.5 9]
	if len(args) == 3 {
		for i := args[0].Value().(int); i < args[1].Value().(int); i += args[2].Value().(int) {
			arr = append(arr, i)
		}
	}

	return compiler.NewNode(arr)
}
