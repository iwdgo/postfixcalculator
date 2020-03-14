package turingmachine

import (
	values "github.com/iwdgo/postfixcalculator/common"
	"fmt"
	"testing"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPNTuringMachine() {
	fmt.Printf("%f", RPNTuringMachine(values.RPNInput))
	// Output: 30.000000
}

func ExampleRPNTuringMachine_exp() {
	fmt.Printf("%.13f", RPNTuringMachine(values.Input))
	// Output: 3.0001220703125
}

func TestRPNTuringMachine(t *testing.T) {
	if got := RPNTuringMachine(values.RPNInput); got != values.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", values.Input, got, values.RPNInputWant)
	}
}

func BenchmarkRPNTuringMachine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNTuringMachine(values.Input)
	}
}
