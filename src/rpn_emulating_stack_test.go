package RPN

import (
	"fmt"
	"testing"
)

/*
Running a single test
*/
func ExampleRPN_emulating_stack() {
	fmt.Printf("%.13f", RPN_emulating_stack(input))
	// Output: 3.0001220703125
}

func ExampleRPN_emulating_stack_sqrt() {
	fmt.Printf("%f", RPN_emulating_stack(RPNInput))
	// Output: 30.000000
}

func TestRPN_emulating_stack(t *testing.T) {
	actual := RPN_emulating_stack(input)
	if actual != 3.0001220703125 {
		t.Errorf("RPN(%s): expected %f, actual %f", input, 3.0001220703125, actual)
	}
}

/*
go test -bench=. allows to test the func
*/
func BenchmarkRPN_emulating_stack(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		RPN_emulating_stack(input)
	}
}
