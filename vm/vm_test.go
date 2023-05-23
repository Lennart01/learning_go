package main

import (
	"testing"

	"github.com/lennart01/learning_go/ast"
)

func TestVM(t *testing.T) {
	// test case 1
	vm1 := NewVM([]Code{
		NewPushCode(1),
		NewPushCode(2),
		NewPlusCode(),
		NewPushCode(3),
		NewMultiplyCode(),
	})
	result1 := vm1.Run()
	if result1.Value().(int) != 9 {
		t.Errorf("Test case 1 failed: expected 9, but got %d", result1.Value().(int))
	}

	// test case 2
	vm2 := NewVM([]Code{
		NewPushCode(1),
		NewPushCode(2),
		NewMultiplyCode(),
		NewPushCode(3),
		NewPlusCode(),
	})
	result2 := vm2.Run()
	if result2.Value().(int) != 5 {
		t.Errorf("Test case 2 failed: expected 5, but got %d", result2.Value().(int))
	}

	// test case 3
	vm3 := NewVM([]Code{
		NewPushCode(1),
		NewPushCode(2),
		NewPlusCode(),
		NewMultiplyCode(),
	})
	result3 := vm3.Run()
	if !result3.IsNothing() {
		t.Errorf("Test case 3 failed: expected Nothing, but got %d", result3.Value().(int))
	}
	// test case 4 (with ast)
	int_exp1 := ast.IntExp{Val: 1}
	int_exp2 := ast.IntExp{Val: 2}
	int_exp3 := ast.IntExp{Val: 3}
	plus_exp := ast.PlusExp{Left: int_exp1, Right: int_exp2}
	mult_exp := ast.MultExp{Left: plus_exp, Right: int_exp3}
	vm4 := LoadAst(mult_exp)
	vm4_result := vm4.Run()
	if vm4_result.Value().(int) != 9 {
		t.Errorf("Test case 4 failed: expected 9, but got %d", vm4_result.Value().(int))
	}

}

func TestLoadAst(t *testing.T) {
	// create an ast
	int_exp1 := ast.IntExp{Val: 1}
	int_exp2 := ast.IntExp{Val: 2}
	plus_exp := ast.PlusExp{Left: int_exp1, Right: int_exp2}
	mult_exp := ast.MultExp{Left: plus_exp, Right: int_exp2}

	// load the ast into the vm
	vm := LoadAst(mult_exp)

	// check that the code was loaded correctly
	expected := []Code{
		NewPushCode(1),
		NewPushCode(2),
		NewPlusCode(),
		NewPushCode(2),
		NewMultiplyCode(),
	}
	for i, code := range vm.code {
		if code != expected[i] {
			t.Errorf("Expected code %v, but got %v", expected[i], code)
		}
	}
}
