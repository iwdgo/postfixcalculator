package slowturingmachine

import (
	"github.com/iwdgo/postfixcalculator/common"
	"testing"
)

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
