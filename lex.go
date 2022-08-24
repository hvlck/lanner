package lanner

import (
	"bufio"
	"math/big"
	"strings"
	"unicode"
)

type Operation int

const (
	ADD = iota
	SUB
	MULT
	DIV
)

type Symbol string
type Number big.Int

type Value interface {
	Symbol | Number | Operation
}

type Node struct {
	Parent *Node
	Kids   []*Node
	// Symbol for now, fix generics situtation later
	Value interface{}
}

func (e *Node) add_node(symbol interface{}) {
	node := MakeNode(symbol)
	node.Parent = e

	e.Kids = append(e.Kids, node)
}

func MakeNode(symbol interface{}) *Node {
	return &Node{
		Kids: []*Node{},
	}
}

func lex(input string) {
	root := &Node{Kids: []*Node{}, Value: "*", Parent: nil}
	reader := bufio.NewReader(strings.NewReader(input))

	for {
		rn, _, err := reader.ReadRune()
		if err != nil {
		}

		switch rn {
		case '+':
			root.add_node(ADD)
		case '-':
			root.add_node(SUB)
		case '*':
			root.add_node(MULT)
		case '/':
			root.add_node(DIV)
		case '(':
		case ')':
		case '>':
		case '<':
		case '=':
		case '!':
		default:
			{
				if unicode.IsLetter(rn) {

				}

				if unicode.IsNumber(rn) {
				}
			}
		}
	}
}
