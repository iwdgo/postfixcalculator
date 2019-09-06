package rpn

import (
	"fmt"
	"testing"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPNTuringMachine() {
	fmt.Printf("%f", RPNTuringMachine(RPNInput))
	// Output: 30.000000
}

func ExampleRPNTuringMachineExp() {
	fmt.Printf("%.13f", RPNTuringMachine(input))
	// Output: 3.0001220703125
}

func TestRPNTuringMachine(t *testing.T) {
	if got := RPNTuringMachine(RPNInput); got != RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", input, got, RPNInputWant)
	}
}

func BenchmarkRPNTuringMachine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNTuringMachine(input)
	}
}
