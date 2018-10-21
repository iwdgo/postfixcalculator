package RPN

import (
	"math"
	"strconv"
	"strings"
	//"fmt"
)

/*

Original is from http://rosettacode.org/wiki/Parsing/RPN_calculator_algorithm#Go where a example of
a printed output can be found. This version contains minor changes for this test.

This emulation is not using a stack in the strict sense as the structure is accessed sequentially
when only pop/push operations are allowed by definition but it is the most efficient as the
"stack" only contains the unused parts of the expression.

*/
func RPN_emulating_stack(input string) float64 {
	var stack []float64 // "stack" only contains numbers
	for _, tok := range strings.Fields(input) {
		switch tok {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "^":
			stack[len(stack)-2] =
				math.Pow(stack[len(stack)-2], stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		case "sqrt":
			stack[len(stack)-1] =
				math.Sqrt(stack[len(stack)-1])
			// Unary operator keeps stack size
		default:
			if f, err := strconv.ParseFloat(tok, 64); err != nil {
				panic("Invalid number and not a known operator")
			} else {
				stack = append(stack, f)
			}
		}
	}
	return stack[0]
}
