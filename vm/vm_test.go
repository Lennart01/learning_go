package main

import (
	"testing"
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
}
