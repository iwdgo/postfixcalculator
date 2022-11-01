package turingmachine

import (
	"fmt"
	values "github.com/iwdgo/postfixcalculator/common"
	"strconv"
	"testing"
)

func ExampleRPNTuringMachine() {
	fmt.Printf("%f", RPNTuringMachine(values.RPNInput))
	// Output: 30.000000
}

func ExampleRPNTuringMachine_exp() {
	fmt.Printf("%.13f", RPNTuringMachine(values.Input))
	// Output: 3.0001220703125
}

func TestRPNTuringMachine(t *testing.T) {
	if got := RPNTuringMachine(values.RPNInput); got != values.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", values.Input, got, values.RPNInputWant)
	}
}

func TestOneOperand(t *testing.T) {
	t.Skip("One operand band is not handled correctly")
	s := "1"
	want, _ := strconv.ParseFloat(s, 64)
	if got := RPNTuringMachine(s); got != want {
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
	RPNTuringMachine(values.InvalidOperator)
}

func TestPanicLeftOperand(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	RPNTuringMachine(values.InvalidLeftOperand)
}

func TestPanicRightOperand(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	RPNTuringMachine(values.InvalidRightOperand)
}

func BenchmarkRPNTuringMachine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNTuringMachine(values.Input)
	}
}
