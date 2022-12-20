package common

const OperatorsList = "+ - * / ^ sqrt"

const (
	RPNInput            = "6 8 4.0 sqrt + 3.01 1.99 + - *" //= 30.000000
	Input               = "3 4 2 * 1 5 - 2 3 ^ ^ / +"      //= 3.0001220703125
	InvalidInput        = "3 4 2 * ? 1 5 - 2 3 ^ ^ / +"    // No value
	InvalidLeftOperand  = "a 4 *"                          // a is not a valid operand
	InvalidRightOperand = "3 4,3 *"                        // 4,3 is not a valid operand
	InvalidOperator     = "3 4 o"                          // o is not a known operator
	OneOperand          = "1"
	OneOperator         = "+"

	RPNInputWant = 30.000000
	InputWant    = 3.0001220703125
)
