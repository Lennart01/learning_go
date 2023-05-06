package main

import (
	"fmt"
	"strconv"
)

type Token int

const (
	EOS Token = iota // End of string
	ZERO
	ONE
	TWO
	OPEN
	CLOSE
	PLUS
	MULT
)

type EXP interface {
	String() string
	Eval() int
}

type Int struct {
	val int
}

func (i *Int) String() string {
	return strconv.Itoa(i.val)
}

func (i *Int) Eval() int {
	return i.val
}

type BinOp struct {
	left  EXP
	right EXP
	op    Token
}

func (b *BinOp) String() string {
	op := ""
	switch b.op {
	case PLUS:
		op = "+"
	case MULT:
		op = "*"
	}
	return fmt.Sprintf("(%s %s %s)", b.left.String(), op, b.right.String())
}
func (b *BinOp) Eval() int {
	switch b.op {
	case PLUS:
		return b.left.Eval() + b.right.Eval()
	case MULT:
		return b.left.Eval() * b.right.Eval()
	default:
		return 0
	}
}

type Parser struct {
	s   string
	pos int
}

func NewParser(s string) *Parser {
	return &Parser{
		s:   s,
		pos: 0,
	}
}

func (p *Parser) skipWhitespace() {
	for p.pos < len(p.s) && (p.s[p.pos] == ' ' || p.s[p.pos] == '\t' || p.s[p.pos] == '\n') {
		p.pos++
	}
}

func (p *Parser) next() Token {
	p.skipWhitespace()
	if p.pos >= len(p.s) {
		return EOS
	}
	switch p.s[p.pos] {
	case '0':
		p.pos++
		return ZERO
	case '1':
		p.pos++
		return ONE
	case '2':
		p.pos++
		return TWO
	case '(':
		p.pos++
		return OPEN
	case ')':
		p.pos++
		return CLOSE
	case '+':
		p.pos++
		return PLUS
	case '*':
		p.pos++
		return MULT
	default:
		return EOS
	}
}

func (p *Parser) parse() EXP {
	return p.parseE()
}

func (p *Parser) parseE() EXP {
	left := p.parseT()
	for {
		switch p.next() {
		case PLUS:
			right := p.parseT()
			left = &BinOp{left, right, PLUS}
		default:
			p.pos--
			return left
		}
	}
}

func (p *Parser) parseT() EXP {
	left := p.parseF()
	for {
		switch p.next() {
		case MULT:
			right := p.parseF()
			left = &BinOp{left, right, MULT}
		default:
			p.pos--
			return left
		}
	}
}

func (p *Parser) parseF() EXP {
	tok := p.next()
	switch tok {
	case ZERO, ONE, TWO:
		val, err := strconv.Atoi(p.s[p.pos-1 : p.pos])
		if err != nil {
			return nil
		}
		return &Int{val}
	case OPEN:
		expr := p.parseE()
		if p.next() == CLOSE {
			return expr
		} else {
			return nil
		}
	default:
		return nil
	}
}

func main() {
	expr := "2 * (1 + 1)"
	parser := NewParser(expr)
	ast := parser.parse()
	fmt.Println("AST:", ast)
	fmt.Println("Result:", ast.Eval())
}
