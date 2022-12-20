package multipleconversions

import (
	"fmt"
	values "github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func ExampleRPN() {
	fmt.Printf("%f", RPN(values.RPNInput))
	// Output: 30.000000
}

func ExampleRPN_exp() {
	fmt.Printf("%.13f", RPN(values.Input))
	// Output: 3.0001220703125
}

func TestRPN(t *testing.T) {
	if got := RPN(values.RPNInput); got != values.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", values.Input, got, values.RPNInputWant)
	}

}

func TestOneOperand(t *testing.T) {
	values.OneOperand(t, RPN)
}

func TestOneOperator(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	RPN("o")
}

func TestPanicOperator(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = RPN(values.InvalidOperator)
}

func BenchmarkRPN(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPN(values.Input)
	}
}
