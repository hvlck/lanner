package ast

import (
	"fmt"
	"strings"

	lanner "github.com/hvlck/lanner/lex"
)

type Node interface {
	Literal() string
	String() string
}

type Expression interface {
	Node
	expression()
}

type Statement interface {
	Node
	statement()
}

type Identifier struct {
	Token lanner.Token
	Value string
}

func (i *Identifier) expression()     {}
func (i *Identifier) Literal() string { return i.Token.Literal }
func (i *Identifier) String() string  { return i.Value }

type PrefixExpression struct {
	Token lanner.Token
	Op    string
	Right Expression
}

func (pe *PrefixExpression) expression()     {}
func (pe *PrefixExpression) Literal() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string  { return "(" + pe.Op + pe.Right.String() + ")" }

type InfixExpression struct {
	Token lanner.Token
	Op    string
	Left  Expression
	Right Expression
}

func (ie *InfixExpression) expression()     {}
func (ie *InfixExpression) Literal() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	return "(" + ie.Left.String() + " " + ie.Op + " " + ie.Right.String() + ")"
}

type NumberLiteral struct {
	Token lanner.Token
	Value int64
}

func (nl *NumberLiteral) expression()     {}
func (nl *NumberLiteral) Literal() string { return nl.Token.Literal }
func (nl *NumberLiteral) String() string  { return fmt.Sprint(nl.Value) }

type BooleanLiteral struct {
	Token lanner.Token
	Value bool
}

func (bl *BooleanLiteral) expression()     {}
func (bl *BooleanLiteral) Literal() string { return bl.Token.Literal }
func (bl *BooleanLiteral) String() string  { return bl.Token.Literal }

type ExpressionStatement struct {
	Token      lanner.Token
	Expression Expression
}

func (es *ExpressionStatement) statement()      {}
func (es *ExpressionStatement) Literal() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string  { return es.Expression.String() }

type Error struct {
	Err   error
	Start lanner.Span
	End   lanner.Span
}

type Program struct {
	Expressions []Expression
	Errors      []Error
}

func (p *Program) String() string {
	var buf strings.Builder

	for _, v := range p.Expressions {
		if v == nil {
			continue
		}
		buf.WriteString(v.String())
	}

	return buf.String()
}

func (p *Program) Literal() string {
	if len(p.Expressions) > 0 {
		return p.Expressions[0].Literal()
	} else {
		return ""
	}
}
