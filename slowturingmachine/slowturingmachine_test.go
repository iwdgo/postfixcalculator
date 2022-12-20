package slowturingmachine

import (
	"fmt"
	values "github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func ExampleRPNSlowTuringMachine() {
	fmt.Printf("%f", RPNSlowTuringMachine(values.RPNInput))
	// Output: 30.000000
}

func ExampleRPNSlowTuringMachine_exp() {
	fmt.Printf("%.13f", RPNSlowTuringMachine(values.Input))
	// Output: 3.0001220703125
}

func TestRPNSlowTuringMachine(t *testing.T) {
	if got := RPNSlowTuringMachine(values.RPNInput); got != values.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", values.RPNInput, got, values.RPNInputWant)
	}
}

func TestPanicLeftOperand(t *testing.T) {
	values.PanicLeftOperand(t, RPNSlowTuringMachine)
}

func TestPanicRightOperand(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = RPNSlowTuringMachine(values.InvalidRightOperand)
}

func TestOneOperand(t *testing.T) {
	t.Skip("one operand is incorrectly handled")
	values.OneOperand(t, RPNSlowTuringMachine)
}

func BenchmarkRPNSlowTuringMachine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNSlowTuringMachine(values.Input)
	}
}
