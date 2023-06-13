package main

import (
	"fmt"
)

func main() {
	// creating test expressions
	int_exp := IntExp{Val: 5}
	int_exp2 := IntExp{Val: 10}
	int_exp3 := IntExp{Val: 15}
	mult_exp := MultExp{Left: int_exp, Right: int_exp}
	plus_exp := PlusExp{Left: int_exp, Right: int_exp}
	plus_exp2 := PlusExp{Left: int_exp2, Right: int_exp3}
	plus_exp3 := PlusExp{Left: plus_exp, Right: plus_exp2}
	plus_exp3 = PlusExp{Left: mult_exp, Right: plus_exp3}
	fmt.Print(plus_exp3.Pretty())
	fmt.Println("=", plus_exp3.Eval())

	int_exp = IntExp{Val: 5}
	int_exp2 = IntExp{Val: 10}
	int_exp3 = IntExp{Val: 15}
	plus_exp = PlusExp{Left: int_exp, Right: int_exp}
	mult_exp = MultExp{Left: plus_exp, Right: int_exp2}
	fmt.Print(mult_exp.Pretty())
	fmt.Println("=", mult_exp.Eval())
}
