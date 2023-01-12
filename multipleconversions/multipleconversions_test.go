package multipleconversions

import (
	"github.com/iwdgo/postfixcalculator/common"
	"testing"
)

func TestOneOperand(t *testing.T) {
	common.OneOperand(t, RPN)
}

func TestOneOperator(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	RPN("o")
}

func TestPanicOperator(t *testing.T) {
	common.PanicOperator(t, RPN)
}

func BenchmarkRPN(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPN(common.Input)
	}
}
