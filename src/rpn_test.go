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

/* */

func TestRPN(t *testing.T) {
	actual := RPN(RPNInput)
	if actual != 30 {
		t.Errorf("RPN(%s): expected %d, actual %f", RPNInput, 30, actual)
	}
}

/*
go test -bench=. allows to test the func main
*/
func BenchmarkRPN(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		RPN(RPNInput)
	}
}
