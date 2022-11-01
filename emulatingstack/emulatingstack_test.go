package emulatingstack

import (
	"fmt"
	values "github.com/iwdgo/postfixcalculator/common"
	"strconv"
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
	t.Skip("One operand band is not handled correctly")
	s := "1"
	want, _ := strconv.ParseFloat(s, 64)
	if got := RPNEmulatingStack(s); got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPanicOperator(t *testing.T) {
	t.Skip("Panics on operand error and not unknown operator")
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	RPNEmulatingStack(values.InvalidOperator)
}

func TestPanicLeftOperand(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	RPNEmulatingStack(values.InvalidLeftOperand)
}

func TestPanicRightOperand(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	RPNEmulatingStack(values.InvalidRightOperand)
}

func BenchmarkRPNEmulatingStack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNEmulatingStack(values.Input)
	}
}
