// pratt_parser.go

package main

import (
	"fmt"
	"strconv"
)

// Token types
const (
	NUMBER = iota
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	LPAREN
	RPAREN
)

// Token represents a token in the input string
type Token struct {
	Type  int
	Value string
}

// Expression represents an expression in the input string
type Expression interface {
	String() string
	Eval() float64
}

// Number represents a numeric value in the input string
type Number struct {
	Value float64
}

// String returns the string representation of the Number
func (n Number) String() string {
	return strconv.FormatFloat(n.Value, 'f', -1, 64)
}

// Eval returns the numeric value of the Number
func (n Number) Eval() float64 {
	return n.Value
}

// BinaryOp represents a binary operation in the input string
type BinaryOp struct {
	Left  Expression
	Right Expression
	Op    int
}

// String returns the string representation of the BinaryOp
func (b BinaryOp) String() string {
	return fmt.Sprintf("(%s %c %s)", b.Left.String(), b.Op, b.Right.String())
}

// Eval returns the numeric value of the BinaryOp
func (b BinaryOp) Eval() float64 {
	switch b.Op {
	case PLUS:
		return b.Left.Eval() + b.Right.Eval()
	case MINUS:
		return b.Left.Eval() - b.Right.Eval()
	case MULTIPLY:
		return b.Left.Eval() * b.Right.Eval()
	case DIVIDE:
		return b.Left.Eval() / b.Right.Eval()
	default:
		return 0
	}
}

// Parser represents a parser for the input string
type Parser struct {
	tokens []Token
	pos    int
}

// NewParser creates a new parser for the given input string
func NewParser(input string) *Parser {
	tokens := tokenize(input)
	return &Parser{tokens: tokens, pos: 0}
}

// parse parses the input string and returns the resulting expression
func (p *Parser) parse() Expression {
	return p.parseExpression(0)
}

// tokenize converts the input string into a list of tokens
func tokenize(input string) []Token {
	var tokens []Token
	i := 0

	for i < len(input) {
		switch input[i] {
		case '+':
			tokens = append(tokens, Token{Type: PLUS, Value: "+"})
			i++
		case '-':
			tokens = append(tokens, Token{Type: MINUS, Value: "-"})
			i++
		case '*':
			tokens = append(tokens, Token{Type: MULTIPLY, Value: "*"})
			i++
		case '/':
			tokens = append(tokens, Token{Type: DIVIDE, Value: "/"})
			i++
		case '(':
			tokens = append(tokens, Token{Type: LPAREN, Value: "("})
			i++
		case ')':
			tokens = append(tokens, Token{Type: RPAREN, Value: ")"})
			i++
		default:
			if isDigit(input[i]) {
				start := i
				for i < len(input) && (isDigit(input[i]) || input[i] == '.') {
					i++
				}
				tokens = append(tokens, Token{Type: NUMBER, Value: input[start:i]})
			} else {
				i++
			}
		}
	}

	return tokens
}

// isDigit returns true if the given character is a digit
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// parseExpression parses an expression with the given precedence
func (p *Parser) parseExpression(precedence int) Expression {
	left := p.parseAtom()

	for p.pos < len(p.tokens) {
		token := p.tokens[p.pos]

		if token.Type != PLUS && token.Type != MINUS && token.Type != MULTIPLY && token.Type != DIVIDE {
			break
		}

		if token.Type == PLUS && precedence <= 1 {
			p.pos++
			right := p.parseExpression(1)
			left = BinaryOp{Left: left, Right: right, Op: PLUS}
		} else if token.Type == MINUS && precedence <= 1 {
			p.pos++
			right := p.parseExpression(1)
			left = BinaryOp{Left: left, Right: right, Op: MINUS}
		} else if token.Type == MULTIPLY && precedence <= 2 {
			p.pos++
			right := p.parseExpression(2)
			left = BinaryOp{Left: left, Right: right, Op: MULTIPLY}
		} else if token.Type == DIVIDE && precedence <= 2 {
			p.pos++
			right := p.parseExpression(2)
			left = BinaryOp{Left: left, Right: right, Op: DIVIDE}
		} else {
			break
		}
	}

	return left
}

// parseAtom parses an atomic expression
func (p *Parser) parseAtom() Expression {
	token := p.tokens[p.pos]

	switch token.Type {
	case NUMBER:
		p.pos++
		value, _ := strconv.ParseFloat(token.Value, 64)
		return Number{Value: value}
	case LPAREN:
		p.pos++
		expr := p.parseExpression(0)
		if p.tokens[p.pos].Type == RPAREN {
			p.pos++
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
