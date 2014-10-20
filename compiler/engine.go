package compiler

import (
	//"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Scanner struct {
	Code   string
	Tokens []string

	unclosed int
	acc      string //accumulator
}

func (scanner *Scanner) emit() {
	if scanner.acc != "" {
		scanner.Tokens = append(scanner.Tokens, scanner.acc)
		scanner.acc = ""
	}
}

func (scanner *Scanner) Tokenize() (tokens []string) {

	for pos, len := 0, len(scanner.Code)-1; pos <= len; pos++ {
		char := byte(scanner.Code[pos])

		if char == '(' {
			scanner.unclosed++
			scanner.emit()
			scanner.acc += string(char)
			scanner.emit()
		} else if char == ')' {
			scanner.unclosed--
			scanner.emit()
			scanner.acc += string(char)
			scanner.emit()
		} else if char == '[' {
			scanner.unclosed++
			scanner.emit()
			scanner.acc += "["
			pos++
			for pos <= len-1 {
				if scanner.Code[pos] == ']' {
					scanner.unclosed--
					scanner.acc += string(scanner.Code[pos])
					break
				}
				scanner.acc += string(scanner.Code[pos])
				pos++
			}
			scanner.emit()

		} else if char == '"' {
			scanner.unclosed++
			scanner.emit()
			scanner.acc += "\""
			pos++
			for pos <= len-1 {
				if scanner.Code[pos] == '"' {
					scanner.unclosed--
					scanner.acc += string(scanner.Code[pos])
					break
				}
				scanner.acc += string(scanner.Code[pos])
				pos++
			}
			scanner.emit()
		} else if char == '\'' {
			scanner.emit()
			scanner.acc += "'"
			pos++
			for pos <= len-1 {
				if scanner.Code[pos] == '(' {
					scanner.unclosed++
				}
				if scanner.Code[pos] == ')' {
					scanner.unclosed--
					scanner.acc += string(scanner.Code[pos])
					break
				}
				scanner.acc += string(scanner.Code[pos])
				pos++
			}
			scanner.emit()

		} else if strings.ContainsAny(string(char), " \r\n\t") {
			scanner.emit()
		} else {
			scanner.acc += string(char)
		}

	}

	scanner.emit()

	if scanner.unclosed != 0 {
		panic("source code unclosed")
	}

	tokens = scanner.Tokens

	return
}

type Parser struct {
	pos    int
	Tokens []string
}

func (parser *Parser) Parse() (nodes []Node) {
	for last := len(parser.Tokens) - 1; parser.pos <= last; parser.pos++ {
		token := parser.Tokens[parser.pos]

		if token == "(" {
			parser.pos++
			nodes = append(nodes, NewNode(parser.Parse()))

		} else if token == ")" {
			break

		} else if token[0] == '\'' { //list
			var list NodeList

			scanner := Scanner{Code: token[1:]}
			listTokens := scanner.Tokenize()
			parser := Parser{Tokens: listTokens}
			listNodes := parser.Parse()

			for _, node := range listNodes[0].Vnode {
				list.PushBack(node.Value())
			}
			nodes = append(nodes, NewNode(list))

		} else if token[0] == '[' { //list
			var arr []interface{}

			scanner := Scanner{Code: token[1 : len(token)-1]}
			listTokens := scanner.Tokenize()

			parser := Parser{Tokens: listTokens}
			listNodes := parser.Parse()

			for _, node := range listNodes {
				arr = append(arr, node.Value())
			}

			//fmt.Println(arr)
			nodes = append(nodes, NewNode(arr))

		} else if token[0] == '"' { //string
			nodes = append(nodes, NewNode(strings.Trim(token, `"`)))

		} else if unicode.IsDigit(int32(token[0])) ||
			(token[0] == '-' && len(token) >= 2 && unicode.IsDigit(int32(token[1]))) {
			if strings.Contains(token, ".") {
				if i, err := strconv.ParseFloat(token, 64); err == nil {
					nodes = append(nodes, NewNode(i))
				}
			} else {
				if i, err := strconv.Atoi(token); err == nil {
					nodes = append(nodes, NewNode(int(i)))
				}
			}

		} else {
			node := Node{
				Type:    Tsymbol,
				Vsymbol: token,
			}
			nodes = append(nodes, node)
		}
	}

	return
}

func Eval(node Node, env Environment) (ret Node) {

	switch node.Type {
	case Tnode:
		{
			if len(node.Vnode) == 0 {
				break
			}
			fn := node.Vnode[0].Vsymbol
			if fn == "" {
				panic("invalid application")
			}

			if build, err := env.Get(fn); err == nil {

				call := build.(Buildin).Func
				ret = call(node.Vnode[1:], env).(Node)
			}
		}
	default:
		if node.Type == Tsymbol {
			if variable, err := env.Get(node.Vsymbol); err == nil {
				ret = variable.(Variable).Value
			}
		} else {
			ret = node
		}
	}

	return
}
