package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"1 + 2 * 0", 1},
		{"2 * (1 + 1)", 4},
		{"(2 + 1) * 0", 0},
		{"1 + (2 * 1) + 2", 5},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			parser := NewParser(tt.input)
			ast := parser.parse()
			got := return_int(ast)
			if got != tt.want {
				t.Errorf("eval(parse(%q)) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func return_int(exp EXP) int {
	switch e := exp.(type) {
	case *Int:
		return e.val
	case *BinOp:
		switch e.op {
		case PLUS:
			return return_int(e.left) + return_int(e.right)
		case MULT:
			return return_int(e.left) * return_int(e.right)
		}
	}
	return 0
}
