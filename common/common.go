package common

import (
	"strconv"
	"testing"
)

const OperatorsList = "+ - * / ^ sqrt"

const (
	RPNInput            = "6 8 4.0 sqrt + 3.01 1.99 + - *" //= 30.000000
	Input               = "3 4 2 * 1 5 - 2 3 ^ ^ / +"      //= 3.0001220703125
	InvalidInput        = "3 4 2 * ? 1 5 - 2 3 ^ ^ / +"    // No value
	InvalidLeftOperand  = "a 4 *"                          // a is not a valid operand
	InvalidRightOperand = "3 4,3 *"                        // 4,3 is not a valid operand
	InvalidOperator     = "3 4 o"                          // o is not a known operator
	OneOperator         = "+"

	RPNInputWant = 30.000000
	InputWant    = 3.0001220703125
)

type Rpn func(s string) float64

func PanicRightOperand(t *testing.T, RPN Rpn) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = RPN(InvalidRightOperand)
}

func PanicLeftOperand(t *testing.T, RPN Rpn) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = RPN(InvalidLeftOperand)
}

func OneInvalidOperator(t *testing.T, RPN Rpn) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = RPN("o")
}

func OneOperand(t *testing.T, RPN Rpn) {
	s := "1"
	want, _ := strconv.ParseFloat(s, 64)
	if got := RPN(s); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func PanicOperator(t *testing.T, RPN Rpn) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = RPN(InvalidOperator)
}
