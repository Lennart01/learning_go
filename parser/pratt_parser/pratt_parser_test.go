package main

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input string
		want  float64
	}{
		{"1 + 2", 3},
		{"1 - 2", -1},
		{"2 * 3", 6},
		{"4 / 2", 2},
		{"(1 + 2) * 3", 9},
		{"2 * (1 + 1 + 1) * 2 + 1", 13},
		{"2 * (1 + 1 + 1) * (2 + 1) - 1 / 2", 17.5},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		got := parser.parse().Eval()
		if got != test.want {
			t.Errorf("parse(%q).Eval() = %v, want %v", test.input, got, test.want)
		}
	}
}
func TestTokenize(t *testing.T) {
	tests := []struct {
		input string
		want  []Token
	}{
		{"1 + 2", []Token{{Type: NUMBER, Value: "1"}, {Type: PLUS, Value: "+"}, {Type: NUMBER, Value: "2"}}},
		{"(1 + 2) * 3", []Token{{Type: LPAREN, Value: "("}, {Type: NUMBER, Value: "1"}, {Type: PLUS, Value: "+"}, {Type: NUMBER, Value: "2"}, {Type: RPAREN, Value: ")"}, {Type: MULTIPLY, Value: "*"}, {Type: NUMBER, Value: "3"}}},
		{"2 * (1 + 1 + 1) * (2 + 1) - 1 / 2", []Token{{Type: NUMBER, Value: "2"}, {Type: MULTIPLY, Value: "*"}, {Type: LPAREN, Value: "("}, {Type: NUMBER, Value: "1"}, {Type: PLUS, Value: "+"}, {Type: NUMBER, Value: "1"}, {Type: PLUS, Value: "+"}, {Type: NUMBER, Value: "1"}, {Type: RPAREN, Value: ")"}, {Type: MULTIPLY, Value: "*"}, {Type: LPAREN, Value: "("}, {Type: NUMBER, Value: "2"}, {Type: PLUS, Value: "+"}, {Type: NUMBER, Value: "1"}, {Type: RPAREN, Value: ")"}, {Type: MINUS, Value: "-"}, {Type: NUMBER, Value: "1"}, {Type: DIVIDE, Value: "/"}, {Type: NUMBER, Value: "2"}}},
		{"1 + 2 * 3", []Token{{Type: NUMBER, Value: "1"}, {Type: PLUS, Value: "+"}, {Type: NUMBER, Value: "2"}, {Type: MULTIPLY, Value: "*"}, {Type: NUMBER, Value: "3"}}},
		{"1 + 2 * 3 - 4 / 5", []Token{{Type: NUMBER, Value: "1"}, {Type: PLUS, Value: "+"}, {Type: NUMBER, Value: "2"}, {Type: MULTIPLY, Value: "*"}, {Type: NUMBER, Value: "3"}, {Type: MINUS, Value: "-"}, {Type: NUMBER, Value: "4"}, {Type: DIVIDE, Value: "/"}, {Type: NUMBER, Value: "5"}}},
	}

	for _, test := range tests {
		got := tokenize(test.input)
		if !tokenSliceEqual(got, test.want) {
			t.Errorf("tokenize(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

// check if two token slices are equal
func tokenSliceEqual(a, b []Token) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Type != b[i].Type || a[i].Value != b[i].Value {
			return false
		}
	}
	return true
}

func TestParseExpression(t *testing.T) {
	tests := []struct {
		input      string
		precedence int
		want       Expression
	}{
		{"1 + 2", 0, BinaryOp{Left: Number{Value: 1}, Right: Number{Value: 2}, Op: PLUS}},
		{"1 + 2 * 3", 0, BinaryOp{Left: Number{Value: 1}, Right: BinaryOp{Left: Number{Value: 2}, Right: Number{Value: 3}, Op: MULTIPLY}, Op: PLUS}},
		{"1 * 2 + 3", 1, BinaryOp{Left: BinaryOp{Left: Number{Value: 1}, Right: Number{Value: 2}, Op: MULTIPLY}, Right: Number{Value: 3}, Op: PLUS}},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		got := parser.parseExpression(test.precedence)
		if !expressionEqual(got, test.want) {
			t.Errorf("parseExpression(%q, %d) = %v, want %v", test.input, test.precedence, got, test.want)
		}
	}
}

// check if two expressions are equal
func expressionEqual(a, b Expression) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.String() != b.String() || a.Eval() != b.Eval() {
		return false
	}
	return true
}
