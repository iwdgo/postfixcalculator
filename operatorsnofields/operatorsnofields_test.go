package operatorsnofields

import (
	values "github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func TestPanicLeftOperand(t *testing.T) {
	values.PanicLeftOperand(t, RPNOperatorsNoFields)
}

func TestPanicRightOperand(t *testing.T) {
	values.PanicRightOperand(t, RPNOperatorsNoFields)
}

func BenchmarkRPNOperatorsNoFields(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNOperatorsNoFields(values.Input)
	}
}
