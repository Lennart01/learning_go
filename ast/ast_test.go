package main

import (
	"testing"
)

func TestAst(t *testing.T) {
	tests := []struct {
		input Exp
		want  int
	}{
		{IntExp{1}, 1},
		{PlusExp{IntExp{1}, IntExp{1}}, 2},
		{MultExp{IntExp{2}, IntExp{2}}, 4},
		{PlusExp{IntExp{1}, MultExp{IntExp{2}, IntExp{2}}}, 5},
		{MultExp{PlusExp{IntExp{1}, IntExp{1}}, IntExp{2}}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.input.Pretty(), func(t *testing.T) {
			got := tt.input.Eval()
			if got != tt.want {
				t.Errorf("eval(%q) = %d, want %d", tt.input.Pretty(), got, tt.want)
			}
		})
	}

}
