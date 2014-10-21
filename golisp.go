package golisp

import (
	"./compiler"
	"./ext"
	//"fmt"
)

var Global compiler.Environment

func Excute(code string) {

	Global.Variables = make(map[string]interface{})

	Global.Register(compiler.NewBuildin("+", compiler.Plus))
	Global.Register(compiler.NewBuildin("-", compiler.Minus))
	Global.Register(compiler.NewBuildin("*", compiler.Multiply))
	Global.Register(compiler.NewBuildin("/", compiler.Divide))
	Global.Register(compiler.NewBuildin("print", compiler.Print))
	Global.Register(compiler.NewBuildin("define", compiler.Define))
	Global.Register(compiler.NewBuildin("apply", compiler.Apply))
	Global.Register(compiler.NewBuildin("lambda", compiler.Lambda))
	Global.Register(compiler.NewBuildin("call", compiler.Call))
	Global.Register(compiler.NewBuildin("map", compiler.Map))
	Global.Register(compiler.NewBuildin("array", compiler.Array))

	Global.Register(compiler.NewBuildin("substr", ext.Substr))
	Global.Register(compiler.NewBuildin("append", ext.Append))
	Global.Register(compiler.NewBuildin("split", ext.Split))
	Global.Register(compiler.NewBuildin("strlen", ext.Strlen))
	Global.Register(compiler.NewBuildin("join", ext.Join))

	compiler.SetGlobal(Global)

	scanner := compiler.Scanner{Code: code}
	tokens := scanner.Tokenize()

	parser := compiler.Parser{Tokens: tokens}
	nodes := parser.Parse()

	for _, node := range nodes {
		compiler.Eval(node, Global)
	}
}
