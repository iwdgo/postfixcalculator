package stack_test

import (
	"fmt"
	"github.com/iwdgo/postfixcalculator/stack"
)

func ExampleRPNStack() {
	fmt.Printf("%f", stack.RPNStack("6 8 4.0 sqrt + 3.01 1.99 + - *"))
	// Output: 30.000000
}

func ExampleRPNStack_exp() {
	fmt.Printf("%.13f", stack.RPNStack("3 4 2 * 1 5 - 2 3 ^ ^ / +"))
	// Output: 3.0001220703125
}
