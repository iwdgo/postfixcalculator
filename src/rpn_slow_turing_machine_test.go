package rpn

import (
	"fmt"
	"testing"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPNSlowTuringMachine() {
	fmt.Printf("%f", RPNSlowTuringMachine(RPNInput))
	// Output: 30.000000
}

func ExampleRPNSlowTuringMachineExp() {
	fmt.Printf("%.13f", RPNSlowTuringMachine(input))
	// Output: 3.0001220703125
}

func TestRPNSlowTuringMachine(t *testing.T) {
	if got := RPNSlowTuringMachine(RPNInput); got != RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", RPNInput, got, RPNInputWant)
	}
}

func BenchmarkRPNSlowTuringMachine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNSlowTuringMachine(input)
	}
}
