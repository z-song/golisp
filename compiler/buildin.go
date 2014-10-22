package compiler

import (
	"fmt"
	"reflect"
)

var Global Environment

func SetGlobal(env Environment) {
	Global = env
}

// (+ 50 50)
func Plus(args []Node, Env Environment) (ret interface{}) {

	first := Eval(args[0], Env)

	if first.Type == Tint {
		sum := first.Value().(int)
		for pos, len := 1, len(args)-1; pos <= len; pos++ {
			sum += Eval(args[pos], Env).ToInt()
		}
		ret = NewNode(sum)
	} else if first.Type == Tdouble {
		sum := first.Value().(float64)
		for pos, len := 2, len(args)-1; pos <= len; pos++ {
			sum += Eval(args[pos], Env).ToDouble()
		}
		ret = NewNode(sum)
	} else {
		sum := first.ToInt()
		for pos, len := 2, len(args)-1; pos <= len; pos++ {
			sum += Eval(args[pos], Env).ToInt()
		}
		ret = NewNode(sum)
	}

	return
}

// (- 100 50)
func Minus(args []Node, Env Environment) (ret interface{}) {
	first := Eval(args[0], Env)
	if first.Type == Tint {
		sum := first.Value().(int)
		for pos, len := 1, len(args)-1; pos <= len; pos++ {
			sum -= Eval(args[pos], Env).ToInt()
		}
		ret = NewNode(sum)
	}
	if first.Type == Tdouble {
		sum := first.Value().(float64)
		for pos, len := 2, len(args)-1; pos <= len; pos++ {
			sum -= Eval(args[pos], Env).ToDouble()
		}
		ret = NewNode(sum)
	}

	return
}

// (* 6 7)
func Multiply(args []Node, Env Environment) (ret interface{}) {
	first := Eval(args[0], Env)
	if first.Type == Tint {
		sum := first.Value().(int)
		for pos, len := 1, len(args)-1; pos <= len; pos++ {
			sum *= Eval(args[pos], Env).ToInt()
		}
		ret = NewNode(sum)
	}
	if first.Type == Tdouble {
		sum := first.Value().(float64)
		for pos, len := 2, len(args)-1; pos <= len; pos++ {
			sum *= Eval(args[pos], Env).ToDouble()
		}
		ret = NewNode(sum)
	}

	return

}

// (/ 10 5)
func Divide(args []Node, Env Environment) (ret interface{}) {
	first := Eval(args[0], Env)
	if first.Type == Tint {
		sum := first.Value().(int)
		for pos, len := 1, len(args)-1; pos <= len; pos++ {
			sum /= Eval(args[pos], Env).ToInt()
		}
		ret = NewNode(sum)
	}
	if first.Type == Tdouble {
		sum := first.Value().(float64)
		for pos, len := 2, len(args)-1; pos <= len; pos++ {
			sum /= Eval(args[pos], Env).ToDouble()
		}
		ret = NewNode(sum)
	}

	return
}

// (print "hello world" (+ 1 2 3 4))
func Print(args []Node, Env Environment) (ret interface{}) {

	for pos, len := 0, len(args)-1; pos <= len; pos++ {

		node := Eval(args[pos], Env)
		fmt.Println(node.Value())
	}
	ret = NewNode(1)

	return
}

// (define test 123)
func Define(args []Node, Env Environment) (ret interface{}) {

	if len(args) != 2 {
		panic("[function define] function need two arguments")
	}

	if args[0].Type != Tsymbol {
		panic("[function define] arguments[1] should be symbol type")
	}

	Env.Register(NewVariable(args[0].Vsymbol, Eval(args[1], Env)))

	return NewNode(0)
}

// (apply '(+ 1 2 3 4))
func Apply(args []Node, Env Environment) (ret interface{}) {
	if len(args) != 2 {
		panic("[function apply] function need two arguments")
	}

	if args[0].Type != Tsymbol {
		panic("[function apply] arguments[1] should be symbol type")
	}

	if args[1].Type != Tlist {
		panic("[function apply] arguments[2] should be list type")
	}

	var nodes []Node
	nodes = append(nodes, args[0])

	for e := args[1].Vlist.Front(); e != nil; e = e.Next() {
		nodes = append(nodes, NewNode(e.Value))
	}

	return Eval(NewNode(nodes), Env)
}

// (lambda (arg...) ...)
func Lambda(args []Node, Env Environment) (ret interface{}) {
	global := Global

	fun := Func{Env: Environment{Parent: &global}}
	for _, arg := range args[0].Vnode {
		fun.Args = append(fun.Args, Arg{Name: arg.Vsymbol})
	}

	fun.Body = args[1].Vnode

	ret = Eval(NewNode(fun), Env)

	return
}

// (call add '(1 2 3 4))
func Call(args []Node, Env Environment) (ret interface{}) {

	op := Eval(args[0], Env)

	if op.Type == Tfunc { //call function

		fun := op.Vfunc
		fun.Env.Variables = make(map[string]interface{})

		for index, _ := range fun.Args {
			fun.Args[index].Value = args[1].Vnode[index]

			fun.Env.Register(NewVariable(fun.Args[index].Name, args[1].Vnode[index]))
		}

		ret = NewNode(Eval(NewNode(fun.Body), fun.Env).Value())
	} else { //call buildin
		ret = (op.Vbuildin)(args[1].Vnode, Env)
	}

	return
}

// (map print '(1 2 3 4))
func Map(args []Node, Env Environment) (ret interface{}) {

	if reflect.TypeOf(args[1].Value()).String() == "compiler.NodeList" { //map function to list
		list := args[1].Value().(NodeList)

		var res NodeList
		for e := list.Front(); e != nil; e = e.Next() {
			nodes := []Node{args[0]}
			nodes = append(nodes, NewNode([]Node{NewNode(e.Value)}))
			result := Call(nodes, Env)

			res.PushBack(result.(Node).Value())
		}

		ret = NewNode(res)

	} else if reflect.TypeOf(args[1].Value()).String() == "[]interface {}" { //map function to array

		var res []interface{}
		for _, ele := range args[1].Value().([]interface{}) {

			nodes := []Node{args[0]}
			nodes = append(nodes, NewNode([]Node{NewNode(ele)}))
			result := Call(nodes, Env)

			res = append(res, result.(Node).Value())
		}

		ret = NewNode(res)
	}

	return
}

// (filter func [3 3 6 7 8])
func Filter(args []Node, Env Environment) (ret interface{}) {
	var res []interface{}
	for _, ele := range args[1].Value().([]interface{}) {

		nodes := []Node{args[0]}
		nodes = append(nodes, NewNode([]Node{NewNode(ele)}))
		result := Call(nodes, Env)

		if result.(Node).ToBool() {
			res = append(res, ele)
		}
	}

	ret = NewNode(res)
	return
}

// todo
func Fold(args []Node, Env Environment) (ret interface{}) {

	return
}

// (array 1 2 3 4 "hello")
func Array(args []Node, Env Environment) (ret interface{}) {

	var res []interface{}
	for _, arg := range args {
		res = append(res, arg.Value())
	}

	return NewNode(res)
}

// (list 1 2 "hi") equals to '(1 2 "hi")
func List(args []Node, Env Environment) (ret interface{}) {
	var list NodeList
	for _, arg := range args {
		list.PushBack(arg.Value())
	}

	return NewNode(list)
}

// (== 1 3)
func Equal(args []Node, Env Environment) (ret interface{}) {
	if len(args) != 2 {
		panic("[function ==] need 2 arguments")
	}

	a, b := Eval(args[0], Env), Eval(args[1], Env)

	is := false
	if a.Type == b.Type {
		is = a.Value() == b.Value()
	}

	ret = NewNode(is)
	return
}

// (!= 2 3)
func Unequal(args []Node, Env Environment) (ret interface{}) {

	if len(args) != 2 {
		panic("[function ==] need 2 arguments")
	}

	a, b := Eval(args[0], Env), Eval(args[1], Env)

	is := true
	if a.Type == b.Type {
		is = a.Value() != b.Value()
	}

	ret = NewNode(is)
	return
}

// (> 2 3)
func Gthan(args []Node, Env Environment) (ret interface{}) {

	if len(args) != 2 {
		panic("[function ==] need 2 arguments")
	}

	a, b := Eval(args[0], Env), Eval(args[1], Env)

	is := false
	if a.Type == b.Type {
		if a.Type == Tint {
			is = a.ToInt() > b.ToInt()
		} else if a.Type == Tdouble {
			is = a.ToDouble() > b.ToDouble()
		} else if a.Type == Tstring {
			is = a.ToString() > b.ToString()
		} else if a.Type == Tbool {
			is = a.ToInt() > b.ToInt()
		}
	}

	ret = NewNode(is)
	return
}

// (< 2 3)
func Lthan(args []Node, Env Environment) (ret interface{}) {

	if len(args) != 2 {
		panic("[function ==] need 2 arguments")
	}

	a, b := Eval(args[0], Env), Eval(args[1], Env)

	is := false
	if a.Type == b.Type {
		if a.Type == Tint {
			is = a.ToInt() < b.ToInt()
		} else if a.Type == Tdouble {
			is = a.ToDouble() < b.ToDouble()
		} else if a.Type == Tstring {
			is = a.ToString() < b.ToString()
		} else if a.Type == Tbool {
			is = a.ToInt() < b.ToInt()
		}
	}

	ret = NewNode(is)
	return
}

// (>= 2 3)
func Gequal(args []Node, Env Environment) (ret interface{}) {

	if len(args) != 2 {
		panic("[function ==] need 2 arguments")
	}

	a, b := Eval(args[0], Env), Eval(args[1], Env)

	is := false
	if a.Type == b.Type {
		if a.Type == Tint {
			is = a.ToInt() >= b.ToInt()
		} else if a.Type == Tdouble {
			is = a.ToDouble() >= b.ToDouble()
		} else if a.Type == Tstring {
			is = a.ToString() >= b.ToString()
		} else if a.Type == Tbool {
			is = a.ToInt() >= b.ToInt()
		}
	}

	ret = NewNode(is)
	return
}

// (<= 2 3)
func Lequal(args []Node, Env Environment) (ret interface{}) {

	if len(args) != 2 {
		panic("[function ==] need 2 arguments")
	}

	a, b := Eval(args[0], Env), Eval(args[1], Env)

	is := false
	if a.Type == b.Type {
		if a.Type == Tint {
			is = a.ToInt() <= b.ToInt()
		} else if a.Type == Tdouble {
			is = a.ToDouble() <= b.ToDouble()
		} else if a.Type == Tstring {
			is = a.ToString() <= b.ToString()
		} else if a.Type == Tbool {
			is = a.ToInt() <= b.ToInt()
		}
	}

	ret = NewNode(is)
	return
}

// todo (&& 2 3)
func And(args []Node, Env Environment) (ret interface{}) {

	if len(args) != 2 {
		panic("[function &&] need 2 arguments")
	}

	a, b := Eval(args[0], Env), Eval(args[1], Env)

	ret = NewNode(a.ToBool() && b.ToBool())

	return
}

// todo (&& 2 3)
func Or(args []Node, Env Environment) (ret interface{}) {

	if len(args) != 2 {
		panic("[function ||] need 2 arguments")
	}

	a, b := Eval(args[0], Env), Eval(args[1], Env)

	ret = NewNode(a.ToBool() || b.ToBool())
	return
}

// (type "hello")
func Type(args []Node, Env Environment) (ret interface{}) {
	if len(args) != 1 {
		panic("[function type] need 1 argument only")
	}

	return NewNode(Eval(args[0], Env).TypeString())
}

//cast
// (string 123)
func String(args []Node, Env Environment) (ret interface{}) {
	if len(args) != 1 {
		panic("[function string] need 1 argument only")
	}

	return NewNode(Eval(args[0], Env).ToString())
}

// (int "123")
func Int(args []Node, Env Environment) (ret interface{}) {

	if len(args) != 1 {
		panic("[function int] need 1 argument only")
	}

	return NewNode(Eval(args[0], Env).ToInt())
}

// (double "123.456")
func Double(args []Node, Env Environment) (ret interface{}) {
	if len(args) != 1 {
		panic("[function double] need 1 argument only")
	}

	return NewNode(Eval(args[0], Env).ToDouble())
}

// (bool "123.456")
func Bool(args []Node, Env Environment) (ret interface{}) {
	if len(args) != 1 {
		panic("[function bool] need 1 argument only")
	}

	return NewNode(Eval(args[0], Env).ToBool())
}

// todo (if (> x y) x y)
func If(args []Node, Env Environment) (ret interface{}) {

	return
}

// todo (if (> x y) x y)
func For(args []Node, Env Environment) (ret interface{}) {

	return
}

// todo (cond (> x y) x y)
func Cond(args []Node, Env Environment) (ret interface{}) {

	return
}

// todo (when (> x y) x y)
func When(args []Node, Env Environment) (ret interface{}) {

	return
}

// todo (when (> x y) x y)
func Car(args []Node, Env Environment) (ret interface{}) {
	list := Eval(args[0], Env)
	if len(args) != 1 {
		panic("[function car] need 1 argument only")
	}
	if list.Type != Tlist {
		panic("[function car] argument 1 should be list type")
	}

	var res NodeList

	res.PushBack(list.Vlist.Front().Value)

	return NewNode(res)
}

// todo (cdr '(1 2 3 4 5))
func Cdr(args []Node, Env Environment) (ret interface{}) {

	list := Eval(args[0], Env)
	if len(args) != 1 {
		panic("[function car] need 1 argument only")
	}
	if list.Type != Tlist {
		panic("[function car] argument 1 should be list type")
	}

	var res NodeList
	for e := list.Vlist.Front().Next(); e != nil; e = e.Next() {
		res.PushBack(e.Value)
	}

	return NewNode(res)
}
