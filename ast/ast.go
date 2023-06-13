package ast

import (
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

// pretty function for plus expression
func (plus_exp PlusExp) Pretty() string {
	return "(" + plus_exp.Left.Pretty() + "+" + plus_exp.Right.Pretty() + ")"
}

// define the mult expression
// implicitly implements the Exp interface
type MultExp struct {
	Left  Exp
	Right Exp
}

// eval function for mult expression
// returns the value of the mult expression
func (mult_exp MultExp) Eval() int {
	return mult_exp.Left.Eval() * mult_exp.Right.Eval()
}

// pretty function for mult expression
func (mult_exp MultExp) Pretty() string {
	return "(" + mult_exp.Left.Pretty() + "*" + mult_exp.Right.Pretty() + ")"
}
