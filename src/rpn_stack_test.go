package RPN

import (
	"fmt"
	"testing"
)

/*
Running exemples where fmt is needed to check output.

*/
func ExampleRPN_stack() {
	fmt.Printf("%.13f", RPN_stack(input))
	// Output: 3.0001220703125
}

func ExampleRPN_stack_sqrt() {
	fmt.Printf("%f", RPN_stack(RPNInput))
	// Output: 30.000000
}

func TestRPN_stack(t *testing.T) {
	if got := RPN_stack(RPNInput); got != RPNInput_want {
		t.Errorf("RPN(%s): got %f, want %f", input, got, RPNInput_want)
	}
}

func BenchmarkRPN_stack(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		RPN_stack(input)
	}
}
