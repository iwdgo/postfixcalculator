package turingmachine_test

import (
	"fmt"
	"github.com/iwdgo/postfixcalculator/turingmachine"
)

func ExampleRPNTuringMachine() {
	fmt.Printf("%f", turingmachine.RPNTuringMachine("6 8 4.0 sqrt + 3.01 1.99 + - *"))
	// Output: 30.000000
}

func ExampleRPNTuringMachine_exp() {
	fmt.Printf("%.13f", turingmachine.RPNTuringMachine("3 4 2 * 1 5 - 2 3 ^ ^ / +"))
	// Output: 3.0001220703125
}
