package multipleconversions

import (
	"fmt"
	"github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func ExampleRPN() {
	fmt.Printf("%f", RPN(common.RPNInput))
	// Output: 30.000000
}

func ExampleRPN_exp() {
	fmt.Printf("%.13f", RPN(common.Input))
	// Output: 3.0001220703125
}

func TestRPN(t *testing.T) {
	if got := RPN(common.RPNInput); got != common.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", common.Input, got, common.RPNInputWant)
	}

}

func TestOneOperand(t *testing.T) {
	common.OneOperand(t, RPN)
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
	common.PanicOperator(t, RPN)
}

func BenchmarkRPN(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPN(common.Input)
	}
}
