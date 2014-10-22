package ext

import (
	"../compiler"
	//"fmt"
	"strings"
)

// (substr "hello world" 1 2)
func Substr(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	if len(args) < 2 || len(args) > 3 {
		panic("[function Substr] function need 3 arguments")
	}

	if args[0].Type != compiler.Tstring || args[1].Type != compiler.Tint || args[2].Type != compiler.Tint {
		panic("[function Substr] invalid arguments")
	}

	str := args[0].Value().(string)
	strlen := int(len(str))
	start := args[1].Value().(int)

	var length int
	if len(args) == 3 {
		length = args[2].Value().(int)
		if length > strlen-start {
			length = strlen - start
		}
	} else {
		length = strlen - start
	}

	if start < 0 {
		start = 0
	}

	if start >= strlen {
		start = strlen
	}

	str = str[start : length+start]

	return compiler.NewNode(str)
}

// (split )
func Split(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	str := args[0].ToString()
	sep := args[1].ToString()

	var res []interface{}
	for _, piece := range strings.Split(str, sep) {
		res = append(res, interface{}(piece))
	}

	return compiler.NewNode(res)
}

// todo (strlen "hwello world")
func Strlen(args []compiler.Node, Env compiler.Environment) (ret interface{}) {
	if args[0].Type != compiler.Tstring {
		panic("[function strlen] argument 1 should be string type")
	}
	if len(args) != 1 {
		panic("[function strlen] need 1 argument only")
	}

	strlen := len(args[0].Vstring)

	return compiler.NewNode(strlen)
}

// todo (join ["hello" "world"] ",")
func Join(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	var strArr []string
	for _, str := range args[0].Varray {
		strArr = append(strArr, str.(string))
	}

	str := strings.Join(strArr, args[1].ToString())

	ret = compiler.NewNode(str)

	return
}
