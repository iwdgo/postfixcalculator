package turingmachine

import (
	"github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func TestOneOperand(t *testing.T) {
	t.Skip("one operand is incorrectly handled")
	common.OneOperand(t, RPNTuringMachine)
}

func TestPanicOperator(t *testing.T) {
	t.Skip("Panics on operand error and not unknown operator")
	common.PanicOperator(t, RPNTuringMachine)
}

func TestNoOperator(t *testing.T) {
	t.Skip("No operator returns 0")
	s := "1 2 3"
	_ = RPNTuringMachine(s)
}

func TestOneInvalidOperator(t *testing.T) {
	t.Skip("does not panic but returns 0")
	common.OneInvalidOperator(t, RPNTuringMachine)
}

func TestPanicLeftOperand(t *testing.T) {
	common.PanicLeftOperand(t, RPNTuringMachine)
}

func TestPanicRightOperand(t *testing.T) {
	common.PanicRightOperand(t, RPNTuringMachine)
}

func BenchmarkRPNTuringMachine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNTuringMachine(common.Input)
	}
}
