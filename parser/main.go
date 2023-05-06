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
}

type Int struct {
	val int
}

func (i *Int) String() string {
	return strconv.Itoa(i.val)
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

func (p *Parser) next() Token {
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
	switch p.next() {
	case ZERO:
		return &Int{0}
	case ONE:
		return &Int{1}
	case TWO:
		return &Int{2}
	case OPEN:
		expr := p.parseE()
		if p.next() == CLOSE {
			return expr
		} else {
			return nil
		}
	default:
		val, err := strconv.Atoi(p.s[p.pos-1 : p.pos])
		if err != nil {
			return nil
		}
		return &Int{val}
	}
}

func main() {
	expr := "1 + 2 * 0"
	parser := NewParser(expr)
	ast := parser.parse()
	fmt.Println(ast)
}
