package golisp

import (
	"./compiler"
	"./ext"
	//"fmt"
)

var Global compiler.Environment

func Excute(code string) {

	Global.Variables = make(map[string]interface{})

	Global.Register(compiler.NewVariable("true", compiler.NewNode(true)))
	Global.Register(compiler.NewVariable("false", compiler.NewNode(false)))

	Global.Register(compiler.NewBuildin("+", compiler.Plus))
	Global.Register(compiler.NewBuildin("-", compiler.Minus))
	Global.Register(compiler.NewBuildin("*", compiler.Multiply))
	Global.Register(compiler.NewBuildin("/", compiler.Divide))
	Global.Register(compiler.NewBuildin("==", compiler.Equal))
	Global.Register(compiler.NewBuildin("!=", compiler.Unequal))
	Global.Register(compiler.NewBuildin(">", compiler.Gthan))
	Global.Register(compiler.NewBuildin("<", compiler.Lthan))
	Global.Register(compiler.NewBuildin(">=", compiler.Gequal))
	Global.Register(compiler.NewBuildin("<=", compiler.Lequal))
	Global.Register(compiler.NewBuildin("&&", compiler.And))
	Global.Register(compiler.NewBuildin("||", compiler.Or))

	Global.Register(compiler.NewBuildin("if", compiler.If))
	Global.Register(compiler.NewBuildin("when", compiler.When))

	Global.Register(compiler.NewBuildin("print", compiler.Print))
	Global.Register(compiler.NewBuildin("define", compiler.Define))
	Global.Register(compiler.NewBuildin("append", compiler.Append))
	Global.Register(compiler.NewBuildin("apply", compiler.Apply))
	Global.Register(compiler.NewBuildin("lambda", compiler.Lambda))
	Global.Register(compiler.NewBuildin("call", compiler.Call))
	Global.Register(compiler.NewBuildin("map", compiler.Map))
	Global.Register(compiler.NewBuildin("filter", compiler.Filter))
	Global.Register(compiler.NewBuildin("array", compiler.Array))
	Global.Register(compiler.NewBuildin("list", compiler.List))

	Global.Register(compiler.NewBuildin("type", compiler.Type))
	Global.Register(compiler.NewBuildin("double", compiler.Double))
	Global.Register(compiler.NewBuildin("int", compiler.Int))
	Global.Register(compiler.NewBuildin("string", compiler.String))
	Global.Register(compiler.NewBuildin("bool", compiler.Bool))

	Global.Register(compiler.NewBuildin("car", compiler.Car))
	Global.Register(compiler.NewBuildin("cdr", compiler.Cdr))

	Global.Register(compiler.NewBuildin("substr", ext.Substr))
	Global.Register(compiler.NewBuildin("split", ext.Split))
	Global.Register(compiler.NewBuildin("strlen", ext.Strlen))
	Global.Register(compiler.NewBuildin("join", ext.Join))

	Global.Register(compiler.NewBuildin("fill", ext.Fill))

	compiler.SetGlobal(Global)

	scanner := compiler.Scanner{Code: code}
	tokens := scanner.Tokenize()

	parser := compiler.Parser{Tokens: tokens}
	nodes := parser.Parse()

	for _, node := range nodes {
		compiler.Eval(node, Global)
	}
}
