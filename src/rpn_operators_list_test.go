package rpn

import (
	"fmt"
	"testing"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPNOperatorsList() {
	fmt.Printf("%f", RPNOperatorsList(RPNInput))
	// Output: 30.000000
}

func ExampleRPNOperatorsListExp() {
	fmt.Printf("%.13f", RPNOperatorsList(input))
	// Output: 3.0001220703125
}

func TestRPNOperatorsList(t *testing.T) {
	if got := RPNOperatorsList(RPNInput); got != RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", RPNInput, got, RPNInputWant)
	}
}

func BenchmarkRPNOperatorsList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNOperatorsList(input)
	}
}
