package RPN

import (
	"fmt"
	"testing"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPN_slow_Turing_machine() {
	fmt.Printf("%f", RPN_slow_Turing_machine(RPNInput))
	// Output: 30.000000
}

func ExampleRPN_slow_Turing_machineExp() {
	fmt.Printf("%.13f", RPN_slow_Turing_machine(input))
	// Output: 3.0001220703125
}

/* */

func TestRPN_slow_Turing_machine(t *testing.T) {
	actual := RPN_slow_Turing_machine(RPNInput)
	if actual != 30 {
		t.Errorf("RPN(%s): expected %d, actual %f", RPNInput, 30, actual)
	}
}

/*
go test -bench=. allows to test the func
*/
func BenchmarkRPN_slow_Turing_machine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPN_slow_Turing_machine(RPNInput)
	}
}
