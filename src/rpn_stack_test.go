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
	actual := RPN_stack(input)
	if actual != 3.0001220703125 {
		t.Errorf("RPN(%s): expected %f, actual %f", input, 3.0001220703125, actual)
	}
}

/*
go test -bench=. allows to test
*/
func BenchmarkRPN_stack(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		RPN_stack(input)
	}
}
