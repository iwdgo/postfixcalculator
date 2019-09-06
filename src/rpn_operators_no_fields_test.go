package rpn

import (
	"fmt"
	"testing"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPN_operators_no_fields() {
	fmt.Printf("%f", RPN_operators_no_fields(RPNInput))
	// Output: 30.000000
}

func ExampleRPN_operators_no_fieldsExp() {
	fmt.Printf("%.13f", RPN_operators_no_fields(input))
	// Output: 3.0001220703125
}

func TestRPN_operators_no_fields(t *testing.T) {
	if got := RPN_operators_no_fields(RPNInput); got != RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", RPNInput, got, RPNInputWant)
	}
}

func BenchmarkRPN_operators_no_fields(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPN_operators_no_fields(input)
	}
}
