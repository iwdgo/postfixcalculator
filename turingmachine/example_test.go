package turingmachine_test

import (
	"fmt"
	values "github.com/iwdgo/postfixcalculator/common"
	"github.com/iwdgo/postfixcalculator/turingmachine"
)

func ExampleRPNTuringMachine() {
	fmt.Printf("%f", turingmachine.RPNTuringMachine(values.RPNInput))
	// Output: 30.000000
}

func ExampleRPNTuringMachine_exp() {
	fmt.Printf("%.13f", turingmachine.RPNTuringMachine(values.Input))
	// Output: 3.0001220703125
}
