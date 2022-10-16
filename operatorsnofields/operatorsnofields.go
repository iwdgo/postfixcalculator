// Package operatorsnofields implements postfix calculator in one loop by editing the string
// containing the expression to evaluate.
package operatorsnofields

import (
	values "github.com/iwdgo/postfixcalculator/common"
	"math"
	"strconv"
	"strings"
)

// RPNOperatorsNoFields returns the result of a sequence of operations written in
// postfix notation (reverse polish notation).
//
// Each operation is replaced by its result until only a number is left.
// It panics if the expression is invalid.
// An invalid sign is interpreted as a value and the next operation panics.
//
// The list of operators is searched using package strings.
// The search avoids to break a second loop but does not provide any major improvement.
func RPNOperatorsNoFields(RPNInput string) float64 {
	words := strings.Fields(RPNInput)
	var num, leftOp, rightOp float64
	var err error
	// Loop until expression contains one word
	for len(words) != 1 {
		for index, w := range words {
			if strings.Contains(values.OperatorsList, w) {
				// w is an operator
				if w == "sqrt" {
					// unary operator
					num, err = strconv.ParseFloat(words[index-1], 64)
					if err != nil {
						panic("Invalid value for sqrt")
					}
					words[index-1] = strconv.FormatFloat(math.Sqrt(num), 'f', 10, 64)
					// Erase operator and operand
					words = append(words[:index], words[index+1:]...)
				} else {
					// Binary operator
					if leftOp, err = strconv.ParseFloat(words[index-2], 64); err != nil {
						panic("Invalid left operand")
					}
					if rightOp, err = strconv.ParseFloat(words[index-1], 64); err != nil {
						panic("Invalid right operand")
					}

					switch w {
					case "+":
						words[index-2] = strconv.FormatFloat(leftOp+rightOp, 'f', 13, 64)
					case "-":
						words[index-2] = strconv.FormatFloat(leftOp-rightOp, 'f', 13, 64)
					case "*":
						words[index-2] = strconv.FormatFloat(leftOp*rightOp, 'f', 13, 64)
					case "/":
						words[index-2] = strconv.FormatFloat(leftOp/rightOp, 'f', 13, 64)
					case "^":
						words[index-2] = strconv.FormatFloat(math.Pow(leftOp, rightOp), 'f', 13, 64)
					default:
						panic("Invalid operator")
					}
					// Erase operator and operands
					words = append(words[:index-1], words[index+1:]...)
				}
				// Restart from start of expression
				break
			}
		}
	}
	num, _ = strconv.ParseFloat(words[0], 64)
	return num
}
