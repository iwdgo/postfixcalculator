package emulatingstack

import (
	"github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func TestRPNEmulatingStack_reservedsign(t *testing.T) {
	common.PanicReservedSign(t, RPNEmulatingStack)
}

func TestRPNEmulatingStack_reservedmark(t *testing.T) {
	common.PanicReservedMark(t, RPNEmulatingStack)
}

func TestOneOperand(t *testing.T) {
	common.OneOperand(t, RPNEmulatingStack)
}

func TestPanicOperator(t *testing.T) {
	t.Skip("Panics on operand error and not unknown operator")
	common.PanicOperator(t, RPNEmulatingStack)
}

func TestPanicLeftOperand(t *testing.T) {
	common.PanicLeftOperand(t, RPNEmulatingStack)
}

func TestPanicRightOperand(t *testing.T) {
	common.PanicRightOperand(t, RPNEmulatingStack)
}

func BenchmarkRPNEmulatingStack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNEmulatingStack(common.Input)
	}
}
