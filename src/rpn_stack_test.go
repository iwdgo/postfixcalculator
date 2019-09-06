package rpn

import (
	"fmt"
	"testing"
)

/*
Running exemples where fmt is needed to check output.

*/
func ExampleRPNStack() {
	fmt.Printf("%.13f", RPNStack(input))
	// Output: 3.0001220703125
}

func ExampleRPNStack_sqrt() {
	fmt.Printf("%f", RPNStack(RPNInput))
	// Output: 30.000000
}

func TestRPNStack(t *testing.T) {
	if got := RPNStack(RPNInput); got != RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", input, got, RPNInputWant)
	}
}

func BenchmarkRPNStack(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		RPNStack(input)
	}
}
