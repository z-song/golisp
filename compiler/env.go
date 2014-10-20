package compiler

import (
	"errors"
	"reflect"
)

type Environment struct {
	Variables map[string]interface{}
	Parent    *Environment
}

func (env *Environment) Register(item interface{}) {

	itemType := reflect.TypeOf(item).String()
	if itemType == "compiler.Buildin" {
		build := item.(Buildin)
		env.Variables[build.Name] = build
	} else if itemType == "compiler.Variable" {
		variable := item.(Variable)
		env.Variables[variable.Name] = variable
	} else {
		panic("item should be variable or function")
	}
}

func (env Environment) Get(name string) (ret interface{}, err error) {

	ok := false

	if ret, ok = env.Variables[name]; !ok {
		if ret, err = env.Parent.Get(name); err != nil {
			err = errors.New("undifined variable or function " + name)
		}
	}

	return
}

type Variable struct {
	Name  string
	Value Node
}

func NewVariable(name string, node Node) Variable {
	return Variable{
		Name:  name,
		Value: node,
	}
}

type Buildin struct {
	Name string
	Func func([]Node, Environment) interface{}
}

func NewBuildin(name string, function func([]Node, Environment) interface{}) Buildin {
	return Buildin{
		Name: name,
		Func: function,
	}
}

type Func struct {
	Name string
	Env  Environment
	Args []Arg
	Body []Node
	Ret  []Node
}

//func (fun Func) String() string {
//	return
//}

type Arg struct {
	Name  string
	Value Node
}
