package stack

import (
	"github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func TestPanicLeftOperand(t *testing.T) {
	common.PanicLeftOperand(t, RPNStack)
}

func TestPanicRightOperand(t *testing.T) {
	common.PanicRightOperand(t, RPNStack)
}

func TestPanicOperator(t *testing.T) {
	common.PanicOperator(t, RPNStack)
}

func BenchmarkRPNStack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNStack(common.Input)
	}
}
