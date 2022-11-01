package operatorslist

import (
	"fmt"
	values "github.com/iwdgo/postfixcalculator/common"
	"testing"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPNOperatorsList() {
	fmt.Printf("%f", RPNOperatorsList(values.RPNInput))
	// Output: 30.000000
}

func ExampleRPNOperatorsList_exp() {
	fmt.Printf("%.13f", RPNOperatorsList(values.Input))
	// Output: 3.0001220703125
}

func TestRPNOperatorsList(t *testing.T) {
	if got := RPNOperatorsList(values.RPNInput); got != values.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", values.RPNInput, got, values.RPNInputWant)
	}
}

func TestPanicLeftOperand(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = RPNOperatorsList(values.InvalidLeftOperand)
}

func TestPanicRightOperand(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = RPNOperatorsList(values.InvalidRightOperand)
}

func BenchmarkRPNOperatorsList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNOperatorsList(values.Input)
	}
}
