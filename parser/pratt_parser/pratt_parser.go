package main

import (
	"fmt"
	"strconv"
)

// Define a set of tokens as constants
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

// Define an interface for expressions
type EXP interface {
	String() string
	Eval() int
}

// Define a struct for integer expressions
type Int struct {
	val int
}

// Implement the String() method for Int expressions
func (i *Int) String() string {
	return strconv.Itoa(i.val)
}

// Implement the Eval() method for Int expressions
func (i *Int) Eval() int {
	return i.val
}

// Define a struct for binary operation expressions
type BinOp struct {
	left  EXP
	right EXP
	op    Token
}

// Implement the String() method for BinOp expressions
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

// Implement the Eval() method for BinOp expressions
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

// Define a struct for the parser
type Parser struct {
	s   string
	pos int
}

// Create a new parser with the given string
func NewParser(s string) *Parser {
	return &Parser{
		s:   s,
		pos: 0,
	}
}

// Skip whitespace characters in the input string
func (p *Parser) skipWhitespace() {
	for p.pos < len(p.s) && (p.s[p.pos] == ' ' || p.s[p.pos] == '\t' || p.s[p.pos] == '\n') {
		p.pos++
	}
}

// Get the next token in the input string
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

// Get the precedence of a token
func (p *Parser) precedence(tok Token) int {
	switch tok {
	case PLUS:
		return 10
	case MULT:
		return 20
	default:
		return -1
	}
}

// Parse an expression
func (p *Parser) parse() EXP {
	return p.parseExpression(0)
}

// Parse an expression with a minimum precedence
func (p *Parser) parseExpression(minPrecedence int) EXP {
	left := p.parsePrimary()

	for {
		tok := p.next()
		precedence := p.precedence(tok)
		if precedence < minPrecedence {
			p.pos--
			return left
		}

		right := p.parseExpression(precedence + 1)
		left = &BinOp{left, right, tok}
	}
}

// Parse a primary expression
func (p *Parser) parsePrimary() EXP {
	tok := p.next()
	switch tok {
	case ZERO, ONE, TWO:
		val, err := strconv.Atoi(p.s[p.pos-1 : p.pos])
		if err != nil {
			return nil
		}
		return &Int{val}
	case OPEN:
		expr := p.parseExpression(0)
		if p.next() == CLOSE {
			return expr
		} else {
			return nil
		}
	default:
		return nil
	}
}

// Main function
func main() {
	expr := "2 * (1 + 1 + 1) * 2 + 1"
	parser := NewParser(expr)
	fmt.Println(expr)
	result := parser.parse()
	fmt.Println(result.String())
	fmt.Println(result.Eval())
}
