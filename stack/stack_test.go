package stack

import (
	"fmt"
	"github.com/iwdgo/postfixcalculator/common"
	"testing"
)

/*
Running exemples where fmt is needed to check output.
*/
func ExampleRPNStack() {
	fmt.Printf("%.13f", RPNStack(common.Input))
	// Output: 3.0001220703125
}

func ExampleRPNStack_sqrt() {
	fmt.Printf("%f", RPNStack(common.RPNInput))
	// Output: 30.000000
}

func TestRPNStack(t *testing.T) {
	if got := RPNStack(common.RPNInput); got != common.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", common.Input, got, common.RPNInputWant)
	}
}

func TestPanicLeftOperand(t *testing.T) {
	common.PanicLeftOperand(t, RPNStack)
}

func TestPanicRightOperand(t *testing.T) {
	common.PanicRightOperand(t, RPNStack)
}

func TestPanicOperator(t *testing.T) {
	common.PanicOperator(t, RPNStack)
}

func BenchmarkRPNStack(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		RPNStack(common.Input)
	}
}
