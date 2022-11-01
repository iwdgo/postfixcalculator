package multipleconversions

import (
	"fmt"
	values "github.com/iwdgo/postfixcalculator/common"
	"strconv"
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
	s := "1"
	want, _ := strconv.ParseFloat(s, 64)
	if got := RPN(s); got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestOneOperator(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	RPN("o")
}

func TestPanicLeftOperand(t *testing.T) {
	t.Skipf("Unreachable. No coverage improvement.")
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = RPN("a sqrt")
}

func TestPanicRightOperand(t *testing.T) {
	t.Skipf("Unreachable. No coverage improvement.")
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = RPN(values.InvalidRightOperand)
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
