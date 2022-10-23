package turingmachine

import (
	"fmt"
	values "github.com/iwdgo/postfixcalculator/common"
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

func TestRPNTuringMachine_oneoperand(t *testing.T) {
	t.Skip("One operand band should return the value")
	i := 1.0
	if got := RPNTuringMachine("1"); got != i {
		t.Fatalf("got %v, want %v", got, i)
	}
}

func TestRPNTuringMachine_panicOperator(t *testing.T) {
	t.Skip("Panics on operand error and not unknown operator")
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	RPNTuringMachine(values.InvalidOperator)
}

func TestRPNTuringMachine_panicLeftOperand(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	RPNTuringMachine(values.InvalidLeftOperand)
}

func TestRPNTuringMachine_panicRightOperand(t *testing.T) {
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
