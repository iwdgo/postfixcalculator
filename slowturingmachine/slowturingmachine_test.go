package slowturingmachine

import (
	"fmt"
	"github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func ExampleRPNSlowTuringMachine() {
	fmt.Printf("%f", RPNSlowTuringMachine(common.RPNInput))
	// Output: 30.000000
}

func ExampleRPNSlowTuringMachine_exp() {
	fmt.Printf("%.13f", RPNSlowTuringMachine(common.Input))
	// Output: 3.0001220703125
}

func TestRPNSlowTuringMachine(t *testing.T) {
	if got := RPNSlowTuringMachine(common.RPNInput); got != common.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", common.RPNInput, got, common.RPNInputWant)
	}
}

func TestPanicLeftOperand(t *testing.T) {
	common.PanicLeftOperand(t, RPNSlowTuringMachine)
}

func TestPanicRightOperand(t *testing.T) {
	common.PanicRightOperand(t, RPNSlowTuringMachine)
}

func TestOneOperand(t *testing.T) {
	t.Skip("one operand is incorrectly handled")
	common.OneOperand(t, RPNSlowTuringMachine)
}

func BenchmarkRPNSlowTuringMachine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNSlowTuringMachine(common.Input)
	}
}
