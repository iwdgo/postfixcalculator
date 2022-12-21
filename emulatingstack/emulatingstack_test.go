package emulatingstack

import (
	"fmt"
	values "github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func ExampleRPNEmulatingStack() {
	fmt.Printf("%f", RPNEmulatingStack(values.RPNInput))
	// Output: 30.000000
}

func ExampleRPNEmulatingStack_exp() {
	fmt.Printf("%.13f", RPNEmulatingStack(values.Input))
	// Output: 3.0001220703125
}

func TestRPNEmulatingStack(t *testing.T) {
	if got := RPNEmulatingStack(values.Input); got != values.InputWant {
		t.Fatalf("RPN(%s): got %f, want %f", values.Input, got, values.InputWant)
	}
}

// To check panic if expression is invalid.
func TestRPNEmulatingStack_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()
	_ = RPNEmulatingStack(values.InvalidInput)
}

func TestOneOperand(t *testing.T) {
	values.OneOperand(t, RPNEmulatingStack)
}

func TestPanicOperator(t *testing.T) {
	t.Skip("Panics on operand error and not unknown operator")
	values.PanicOperator(t, RPNEmulatingStack)
}

func TestPanicLeftOperand(t *testing.T) {
	values.PanicLeftOperand(t, RPNEmulatingStack)
}

func TestPanicRightOperand(t *testing.T) {
	values.PanicRightOperand(t, RPNEmulatingStack)
}

func BenchmarkRPNEmulatingStack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNEmulatingStack(values.Input)
	}
}
