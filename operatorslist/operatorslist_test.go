package operatorslist

import (
	"github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func TestPanicLeftOperand(t *testing.T) {
	common.PanicLeftOperand(t, RPNOperatorsList)
}

func TestPanicRightOperand(t *testing.T) {
	common.PanicRightOperand(t, RPNOperatorsList)
}

func BenchmarkRPNOperatorsList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNOperatorsList(common.Input)
	}
}
