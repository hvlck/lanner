package parse

import (
	"strconv"

	"github.com/hvlck/lanner/ast"
	lanner "github.com/hvlck/lanner/lex"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LTGT
	LTGTEQ
	SUM
	PRODUCT
	PREFIX
	CALL
)

var prec_table = map[lanner.TokenType]int{
	lanner.EQUALITY:     EQUALS,
	lanner.NOT_EQUALITY: EQUALS,
	lanner.LT:           LTGT,
	lanner.GT:           LTGT,
	lanner.LTEQ:         LTGTEQ,
	lanner.GTEQ:         LTGTEQ,
	lanner.ADD:          SUM,
	lanner.SUB:          SUM,
	lanner.MULT:         PRODUCT,
	lanner.DIV:          PRODUCT,
	lanner.LPAREN:       CALL,
}

func prec(tok lanner.Token) int {
	if precedence, ok := prec_table[tok.Type]; ok {
		return precedence
	}

	return LOWEST
}

type Parser struct {
	lexer *lanner.Lexer

	current, next lanner.Token

	prefixes map[lanner.TokenType]func() ast.Expression
	infixes  map[lanner.TokenType]func(ast.Expression) ast.Expression
}

func (p *Parser) nextToken() {
	p.current = p.next
	p.next = p.lexer.Next()
}

func (p *Parser) nextTokenIs(t lanner.TokenType) bool {
	return p.next.Type == t
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	pref := p.prefixes[p.current.Type]
	if pref == nil {
		return nil
	}

	left := pref()
	for !p.nextTokenIs(lanner.LN_BREAK) && precedence < prec(p.next) {
		infix := p.infixes[p.next.Type]
		if infix == nil {
			return left
		}

		p.nextToken()
		left = infix(left)
	}

	return left
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Token: p.current,
		Op:    p.current.Literal,
		Left:  left,
	}

	precedence := prec(p.current)
	p.nextToken()
	expr.Right = p.parseExpression(precedence)

	return expr
}

func (p *Parser) parseNumberLiteral() ast.Expression {
	lit := &ast.NumberLiteral{Token: p.current}
	val, err := strconv.ParseInt(p.current.Literal, 0, 64)
	if err != nil {
		return nil
	}

	lit.Value = val
	return lit
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{
		Expressions: []ast.Expression{},
	}

	for p.current.Type != lanner.EOI {
		stmt := p.parseExpression(LOWEST)
		if stmt != nil {
			program.Expressions = append(program.Expressions, stmt)
		}

		p.nextToken()
	}

	return program
}

func New(lexer *lanner.Lexer) *Parser {
	parser := &Parser{
		lexer: lexer,
	}

	parser.prefixes = map[lanner.TokenType]func() ast.Expression{
		lanner.INT: parser.parseNumberLiteral,
	}

	parser.infixes = map[lanner.TokenType]func(ast.Expression) ast.Expression{
		lanner.ADD:          parser.parseInfixExpression,
		lanner.SUB:          parser.parseInfixExpression,
		lanner.DIV:          parser.parseInfixExpression,
		lanner.MULT:         parser.parseInfixExpression,
		lanner.EQUALITY:     parser.parseInfixExpression,
		lanner.NOT_EQUALITY: parser.parseInfixExpression,
		lanner.GT:           parser.parseInfixExpression,
		lanner.LT:           parser.parseInfixExpression,
		lanner.GTEQ:         parser.parseInfixExpression,
		lanner.LTEQ:         parser.parseInfixExpression,
	}

	parser.nextToken()
	parser.nextToken()

	return parser
}
