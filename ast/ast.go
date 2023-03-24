package main

import (
	"fmt"
	"strconv"
)

// define the base interface that all expressions implement
type Exp interface {
	Eval() int
	Pretty() string
}

// define the int expression
// implicitly implements the Exp interface
type IntExp struct {
	Val int
}

// eval function for int expression
// returns the value of the int expression
func (int_exp IntExp) Eval() int {
	return int_exp.Val
}

// pretty function for int expression
func (int_exp IntExp) Pretty() string {
	return strconv.Itoa(int_exp.Val)
}

// define the plus expression
// implicitly implements the Exp interface
type PlusExp struct {
	Left  Exp
	Right Exp
}

// eval function for plus expression
// returns the value of the plus expression
func (plus_exp PlusExp) Eval() int {
	return plus_exp.Left.Eval() + plus_exp.Right.Eval()
}

func (plus_exp PlusExp) Pretty() string {
	return "(" + plus_exp.Left.Pretty() + "+" + plus_exp.Right.Pretty() + ")"
}

func main() {
	// create an int expression
	int_exp := IntExp{Val: 5}
	int_exp2 := IntExp{Val: 10}
	int_exp3 := IntExp{Val: 15}

	// create a plus expression
	plus_exp := PlusExp{Left: int_exp, Right: int_exp}
	plus_exp2 := PlusExp{Left: int_exp2, Right: int_exp3}
	plus_exp3 := PlusExp{Left: plus_exp, Right: plus_exp2}
	fmt.Print(plus_exp3.Pretty())

}
