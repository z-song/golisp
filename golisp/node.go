package golisp

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
	Tnode
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

	Vstring string
	Vint    int64
	Vdouble float64
	Vbool   bool
	Vlist   NodeList
	Vnode   []Node
	Vsymbol string
	Vfunc   Func
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

	default:
		ret = nil
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
			Vint: data.(int64),
		}
	} else if kind == reflect.Int {
		node = Node{
			Type: Tint,
			Vint: int64(data.(int)),
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
	} else if reflect.TypeOf(data).String() == "golisp.NodeList" {
		node = Node{
			Type:  Tlist,
			Vlist: data.(NodeList),
		}
	} else if reflect.TypeOf(data).String() == "golisp.Func" {
		node = Node{
			Type:  Tfunc,
			Vfunc: data.(Func),
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

func (node Node) ToInt() (ret int64, err error) {

	ret = 0

	if node.Type == Tint {
		ret = node.Vint
	}
	if node.Type == Tdouble {
		ret = int64(node.Vdouble)
	}
	if node.Type == Tstring {
		tmp, _ := strconv.Atoi(node.Vstring)
		ret = int64(tmp)
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

func (node Node) ToDouble() (ret float64, err error) {
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
