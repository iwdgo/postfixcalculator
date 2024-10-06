package turingmachine

import (
	"github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func TestOneOperand(t *testing.T) {
	common.OneOperand(t, RPNTuringMachine)
}

func TestPanicOperator(t *testing.T) {
	common.PanicOperator(t, RPNTuringMachine)
}

func TestNoOperator(t *testing.T) {
	s := "1 2 3"
	_ = RPNTuringMachine(s)
}

func TestOneInvalidOperator(t *testing.T) {
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
