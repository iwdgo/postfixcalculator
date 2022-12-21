package emulatingstack

import (
	"fmt"
	"github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func ExampleRPNEmulatingStack() {
	fmt.Printf("%f", RPNEmulatingStack(common.RPNInput))
	// Output: 30.000000
}

func ExampleRPNEmulatingStack_exp() {
	fmt.Printf("%.13f", RPNEmulatingStack(common.Input))
	// Output: 3.0001220703125
}

func TestRPNEmulatingStack(t *testing.T) {
	if got := RPNEmulatingStack(common.Input); got != common.InputWant {
		t.Fatalf("RPN(%s): got %f, want %f", common.Input, got, common.InputWant)
	}
}

// To check panic if expression is invalid.
func TestRPNEmulatingStack_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()
	_ = RPNEmulatingStack(common.InvalidInput)
}

func TestOneOperand(t *testing.T) {
	common.OneOperand(t, RPNEmulatingStack)
}

func TestPanicOperator(t *testing.T) {
	t.Skip("Panics on operand error and not unknown operator")
	common.PanicOperator(t, RPNEmulatingStack)
}

func TestPanicLeftOperand(t *testing.T) {
	common.PanicLeftOperand(t, RPNEmulatingStack)
}

func TestPanicRightOperand(t *testing.T) {
	common.PanicRightOperand(t, RPNEmulatingStack)
}

func BenchmarkRPNEmulatingStack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNEmulatingStack(common.Input)
	}
}
