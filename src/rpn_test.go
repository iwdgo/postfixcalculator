package RPN

import (
	"fmt"
	"testing"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPN() {
	fmt.Printf("%f", RPN(RPNInput))
	// Output: 30.000000
}

func ExampleRPNExp() {
	fmt.Printf("%.13f", RPN(input))
	// Output: 3.0001220703125
}

func TestRPN(t *testing.T) {
	if got := RPN(RPNInput); got != RPNInput_want {
		t.Errorf("RPN(%s): got %f, want %f", input, got, RPNInput_want)
	}

}

func BenchmarkRPN(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPN(input)
	}
}
