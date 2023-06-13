package main

import (
	"container/list"
	"strconv"

	"github.com/lennart01/learning_go/ast.go"
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

func (vm *VM) transformAst(ast_exp ast.Exp) {
	// switch case on the type of the ast
	switch ast_exp := ast_exp.(type) {
	// if the ast is an int expression
	case ast.IntExp:
		// push the value onto the stack
		vm.code = append(vm.code, NewPushCode(ast_exp.Eval()))
	// if the ast is a plus expression
	case ast.PlusExp:
		// parse the left and right expressions
		vm.transformAst(ast_exp.Left)
		vm.transformAst(ast_exp.Right)
		// push a plus code onto the stack
		vm.code = append(vm.code, NewPlusCode())
	// if the ast is a mult expression
	case ast.MultExp:
		// parse the left and right expressions
		vm.transformAst(ast_exp.Left)
		vm.transformAst(ast_exp.Right)
		// push a multiply code onto the stack
		vm.code = append(vm.code, NewMultiplyCode())
	}
}

// loads an ast into the vm
func LoadAst(ast ast.Exp) VM {
	// create a new vm
	vm := NewVM([]Code{})
	// parse the ast into code
	vm.transformAst(ast)
	return vm
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
			// otherwise, push the sum of the two values onto the stack
			if vm.stack.Len() < 2 {
				return Nothing()
			}
			val1 := vm.stack.Remove(vm.stack.Back()).(int)
			val2 := vm.stack.Remove(vm.stack.Back()).(int)
			vm.stack.PushBack(val1 + val2)
		case MULTIPLY:
			// pop the top two values
			// if the stack is empty, return Nothing
			// otherwise, push the product of the two values onto the stack
			if vm.stack.Len() < 2 {
				return Nothing()
			}
			val1 := vm.stack.Remove(vm.stack.Back()).(int)
			val2 := vm.stack.Remove(vm.stack.Back()).(int)
			vm.stack.PushBack(val1 * val2)
		}
	}
	// if the stack is empty, return Nothing
	if vm.stack.Len() == 0 {
		return Nothing()
	}
	// otherwise, return the top value
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

// prints the calculation
func showCalculation(vm VM) {
	stack_list := list.New()
	prev_value := ""
	// loop through the code
	// "_" ignores the index
	for _, code := range vm.code {
		// switch case on the opcode
		switch code.Op {
		case PUSH:
			// push the value onto the stack_list
			stack_list.PushBack(code.val)
		case PLUS:
			// pop the top two values and set the prev_value to the calculation
			if prev_value == "" {
				element1 := stack_list.Back()
				element2 := stack_list.Back().Prev()
				val1 := strconv.Itoa(stack_list.Remove(element1).(int))
				val2 := strconv.Itoa(stack_list.Remove(element2).(int))
				prev_value = "(" + val1 + " + " + val2 + ")"
				// append the new calculation to the prev_value
			} else if prev_value != "" {
				element2 := stack_list.Back()
				val2 := strconv.Itoa(stack_list.Remove(element2).(int))
				val1 := prev_value
				prev_value = "(" + val1 + " + " + val2 + ")"
			}
		// pop the top two values and set the prev_value to the calculation
		case MULTIPLY:
			if prev_value == "" {
				element1 := stack_list.Back()
				element2 := stack_list.Back().Prev()
				val1 := strconv.Itoa(stack_list.Remove(element1).(int))
				val2 := strconv.Itoa(stack_list.Remove(element2).(int))
				prev_value = "(" + val1 + " * " + val2 + ")"
				// append the new calculation to the prev_value
			} else if prev_value != "" {
				element2 := stack_list.Back()
				val2 := strconv.Itoa(stack_list.Remove(element2).(int))
				val1 := prev_value
				prev_value = "(" + val1 + " * " + val2 + ")"
			}
		}
	}
	// print the calculation
	print("The VM runs the following calculation: \n")
	print(prev_value)
	println()
}

// test the vm
func main() {
	// create a new vm with the code
	vm := NewVM([]Code{
		NewPushCode(1),
		NewPushCode(2),
		NewPlusCode(),
		NewPushCode(3),
		NewMultiplyCode(),
	})
	// print the calculation
	showCalculation(vm)
	// run the vm
	result := vm.Run()
	// print the result
	showVMResult(result)

	// create an ast
	int_exp1 := ast.IntExp{Val: 1}
	int_exp2 := ast.IntExp{Val: 2}
	plus_exp := ast.PlusExp{Left: int_exp1, Right: int_exp2}
	mult_exp := ast.MultExp{Left: plus_exp, Right: int_exp2}
	// load the ast into the vm
	vm2 := LoadAst(mult_exp)
	// print the calculation
	showCalculation(vm2)
	// run the vm
	result2 := vm2.Run()
	// print the result
	showVMResult(result2)
}
