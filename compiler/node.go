package compiler

import (
	"container/list"
	"fmt"
	"reflect"
	"strconv"
)

type NodeType int

const (
	Tnil NodeType = iota
	Tint
	Tdouble
	Tbool
	Tstring
	Tlist
	Tsymbol
	Tfunc
	Tbuildin
	Tnode
	Tarray
)

type NodeList struct {
	list.List
}

func (l NodeList) String() (ret string) {
	ret = "'("
	for e := l.Front(); e != nil; e = e.Next() {
		ret += fmt.Sprintf("%v ", e.Value)
	}
	ret = ret[:len(ret)-1] + ")"

	return
}

type Node struct {
	Type NodeType

	Vstring  string
	Vint     int
	Vdouble  float64
	Vbool    bool
	Vlist    NodeList
	Vnode    []Node
	Vsymbol  string
	Vfunc    Func
	Vbuildin func([]Node, Environment) interface{}
	Varray   []interface{}
}

func (node Node) Value() (ret interface{}) {
	switch node.Type {
	case Tnil:
		ret = "nil"
	case Tint:
		ret = node.Vint
	case Tdouble:
		ret = node.Vdouble
	case Tbool:
		ret = node.Vbool
	case Tstring:
		ret = node.Vstring
	case Tlist:
		ret = node.Vlist
	case Tnode:
		ret = node.Vnode
	case Tsymbol:
		ret = node.Vsymbol
	case Tfunc:

		format := "lambda("
		for _, arg := range node.Vfunc.Args {
			format += arg.Name + " "
		}
		format = format[:len(format)-1] + ")"
		ret = format
	case Tarray:
		ret = node.Varray
	default:
		ret = nil
	}

	return
}

func (node Node) TypeString() (ret string) {
	switch node.Type {
	case Tnil:
		ret = "nil"
	case Tint:
		ret = "int"
	case Tdouble:
		ret = "double"
	case Tbool:
		ret = "bool"
	case Tstring:
		ret = "string"
	case Tlist:
		ret = "list"
	case Tnode:
		ret = "nodes"
	case Tsymbol:
		ret = "symble"
	case Tfunc:
		ret = "lambda"
	case Tarray:
		ret = "array"
	default:
		ret = "unknown"
	}

	return
}

func (node Node) String() (ret string) {
	switch node.Type {
	case Tnil:
		ret = "[nil] : nil"
	case Tint:
		ret = "[int] : " + strconv.Itoa(int(node.Vint))
	case Tdouble:
		ret = "[double] : " + strconv.FormatFloat(node.Vdouble, 'f', 6, 64)
	case Tbool:
		ret = "[bool] : "
		if node.Vbool {
			ret += "true"
		} else {
			ret += "false"
		}
	case Tstring:
		ret = "[string] : " + node.Vstring
	case Tlist:
		ret = node.Vlist.String()
	case Tnode:
		ret = "[nodes] : " //node.Vnode
		for _, item := range node.Vnode {
			ret += item.String() + "\n"
		}

	case Tsymbol:
		ret = "[symbol] : " + node.Vsymbol
	case Tfunc:
		format := "lambda("
		for _, arg := range node.Vfunc.Args {
			format += arg.Name + " "
		}
		format = format[:len(format)-1] + ")"
		ret = format
	case Tarray:
		ret = "[array] : ["
		for _, item := range node.Varray {
			ret += NewNode(item).ToString() + " "
		}
		ret = ret[:len(ret)-1] + "]"
	case Tbuildin:
		ret = "buildin"
	}

	return
}

func NewNode(data interface{}) (node Node) {

	kind := reflect.TypeOf(data).Kind()

	if kind == reflect.String {
		node = Node{
			Type:    Tstring,
			Vstring: data.(string),
		}
	} else if kind == reflect.Int64 {
		node = Node{
			Type: Tint,
			Vint: data.(int),
		}
	} else if kind == reflect.Int {
		node = Node{
			Type: Tint,
			Vint: data.(int),
		}
	} else if kind == reflect.Float64 {
		node = Node{
			Type:    Tdouble,
			Vdouble: data.(float64),
		}
	} else if kind == reflect.Bool {
		node = Node{
			Type:  Tbool,
			Vbool: data.(bool),
		}
	} else if reflect.TypeOf(data).String() == "compiler.NodeList" {
		node = Node{
			Type:  Tlist,
			Vlist: data.(NodeList),
		}
	} else if reflect.TypeOf(data).String() == "compiler.Func" {
		node = Node{
			Type:  Tfunc,
			Vfunc: data.(Func),
		}

	} else if reflect.TypeOf(data).String() == "func([]compiler.Node, compiler.Environment) interface {}" {
		node = Node{
			Type:     Tbuildin,
			Vbuildin: data.(func([]Node, Environment) interface{}),
		}

	} else if reflect.TypeOf(data).String() == "[]interface {}" {
		node = Node{
			Type:   Tarray,
			Varray: data.([]interface{}),
		}
	} else if kind == reflect.Slice {
		node = Node{
			Type:  Tnode,
			Vnode: data.([]Node),
		}
	} else {
		node = Node{
			Type: Tnil,
		}
	}

	return
}

func (node Node) ToString() (ret string) {
	if node.Type == Tint {
		ret = strconv.Itoa(node.Vint)
	} else if node.Type == Tdouble {
		ret = strconv.FormatFloat(node.Vdouble, 'f', 6, 64)
	} else if node.Type == Tstring {
		ret = node.Vstring
	} else if node.Type == Tarray {

		for _, item := range node.Varray {
			ret += NewNode(item).ToString()
		}

	} else if node.Type == Tbool {
		if node.Vbool {
			ret = "true"
		} else {
			ret = "false"
		}
	}

	return
}

func (node Node) ToInt() (ret int) {

	ret = 0

	if node.Type == Tint {
		ret = node.Vint
	}
	if node.Type == Tdouble {
		ret = int(node.Vdouble)
	}
	if node.Type == Tstring {
		tmp, _ := strconv.Atoi(node.Vstring)
		ret = int(tmp)
	}
	if node.Type == Tbool {
		if node.Vbool {
			ret = 1
		} else {
			ret = 0
		}
	}

	return
}

func (node Node) ToDouble() (ret float64) {
	ret = 0
	if node.Type == Tint {
		ret = float64(node.Vint)
	}
	if node.Type == Tdouble {
		ret = node.Vdouble
	}
	if node.Type == Tstring {
		ret, _ = strconv.ParseFloat(node.Vstring, 64)

	}
	if node.Type == Tbool {
		if node.Vbool {
			ret = 1
		} else {
			ret = 0
		}
	}

	return
}

func (node Node) ToBool() (ret bool) {
	ret = false
	if node.Type == Tint {
		ret = node.Vint != 0
	}
	if node.Type == Tdouble {
		ret = node.Vdouble != 0
	}
	if node.Type == Tstring {
		ret = node.Vstring != ""

	}
	if node.Type == Tbool {
		ret = node.Vbool
	}
	if node.Type == Tarray {
		ret = len(node.Varray) != 0
	}

	return
}
