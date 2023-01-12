package slowturingmachine_test

import (
	"fmt"
	"github.com/iwdgo/postfixcalculator/slowturingmachine"
)

func ExampleRPNSlowTuringMachine() {
	fmt.Printf("%f", slowturingmachine.RPNSlowTuringMachine("6 8 4.0 sqrt + 3.01 1.99 + - *"))
	// Output: 30.000000
}

func ExampleRPNSlowTuringMachine_exp() {
	fmt.Printf("%.13f", slowturingmachine.RPNSlowTuringMachine("3 4 2 * 1 5 - 2 3 ^ ^ / +"))
	// Output: 3.0001220703125
}
