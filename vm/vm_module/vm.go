package main

import (
	"container/list"
)

// Optional type is used to represent a value that may or may not exist
type Optional struct {
	val    interface{} // can be any type
	exists bool        // true if the value exists
}

// helper functions for Optional
// returns true if the value exists
func (o Optional) IsJust() bool {
	return o.exists
}

// returns the value if it exists
// if it does not exist panic
func (o Optional) Value() interface{} {
	if !o.exists {
		panic("Value does not exist")
	}
	return o.val
}

// creates a new Optional with the value
func Just(val interface{}) Optional {
	return Optional{val, true}
}

// creates a new Optional with no value
func Nothing() Optional {
	return Optional{nil, false}
}

// returns true if the value does not exist
func (o Optional) IsNothing() bool {
	return !o.exists
}

// type defines a new type similar to typedef in c++
type OpCode int

// define constants for the opcodes
const (
	PUSH OpCode = iota
	PLUS
	MULTIPLY
)

// define a struct to represent a code
type Code struct {
	Op  OpCode
	val int
}

// helper functions for Code
// these functions are used to push a code onto the stack
func NewPushCode(val int) Code {
	return Code{PUSH, val}
}
func NewPlusCode() Code {
	return Code{PLUS, 0}
}
func NewMultiplyCode() Code {
	return Code{MULTIPLY, 0}
}

// define a struct to represent a virtual machine
type VM struct {
	code  []Code     // holds the program code
	stack *list.List // holds the stack
}

// Creates a new vm
func NewVM(code []Code) VM {
	return VM{code, list.New()} // initialize the stack as an empty list
}

// Runs the program
func (vm VM) Run() Optional {
	// always start with an empty stack
	vm.stack.Init()

	// loop through the code
	// "_" ignores the index
	for _, code := range vm.code {
		// switch case on the opcode
		switch code.Op {
		// push the value onto the stack
		case PUSH:
			vm.stack.PushBack(code.val)
		case PLUS:
			// pop the top two values
			// if the stack is empty, return Nothing
			// otherwise, return the sum of the two values
			if vm.stack.Len() < 2 {
				return Nothing()
			}
			val1 := vm.stack.Remove(vm.stack.Back()).(int)
			val2 := vm.stack.Remove(vm.stack.Back()).(int)
			vm.stack.PushBack(val1 + val2)
		case MULTIPLY:
			// pop the top two values
			// if the stack is empty, return Nothing
			// otherwise, return the product of the two values
			if vm.stack.Len() < 2 {
				return Nothing()
			}
			val1 := vm.stack.Remove(vm.stack.Back()).(int)
			val2 := vm.stack.Remove(vm.stack.Back()).(int)
			vm.stack.PushBack(val1 * val2)
		}
	}
	if vm.stack.Len() == 0 {
		return Nothing()
	}
	return Just(vm.stack.Back().Value.(int))
}

// prints the result of the vm
func showVMResult(result Optional) {
	if result.IsJust() {
		println(result.Value().(int))
	} else {
		println("Nothing")
	}

}
