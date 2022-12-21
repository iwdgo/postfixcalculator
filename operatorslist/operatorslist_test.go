package operatorslist

import (
	"fmt"
	"github.com/iwdgo/postfixcalculator/common"
	"testing"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPNOperatorsList() {
	fmt.Printf("%f", RPNOperatorsList(common.RPNInput))
	// Output: 30.000000
}

func ExampleRPNOperatorsList_exp() {
	fmt.Printf("%.13f", RPNOperatorsList(common.Input))
	// Output: 3.0001220703125
}

func TestRPNOperatorsList(t *testing.T) {
	if got := RPNOperatorsList(common.RPNInput); got != common.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", common.RPNInput, got, common.RPNInputWant)
	}
}

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
