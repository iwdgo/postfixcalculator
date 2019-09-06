package rpn

import (
	"math"
	"strconv"
	"strings"
)

/*

Returns the result of a string in reverse polish notation (postfix) by exploding the string
to a slice and editing the slice by replacing each op by its result until only a number is left or
failing if the expression is invalid.

Searching the operators list as the constant string using the strings. package
It removes breaking a second loop without major improvement.

An invalid sign is interpreted as a value and the next operation panics.

*/
/* */
func RPN_operators_no_fields(RPNInput string) float64 {
	words := strings.Fields(RPNInput)
	num := 0.0
	leftOp := 0.0
	rightOp := 0.0
	for len(words) != 1 { //length of expression stops processing
		err = nil
		for index, w := range words {
			if strings.Contains(operatorsList, w) {
				//w is an operator
				if w == "sqrt" {
					//unary operator
					num, err = strconv.ParseFloat(words[index-1], 64)
					if err != nil {
						panic("Invalid value for sqrt")
					}
					words[index-1] = strconv.FormatFloat(math.Sqrt(num), 'f', 10, 64)
					words = append(words[:index], words[index+1:]...) //removing sqrt
				} else {
					//binary operator
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
					/* removing binary operator by removing the slice until the item after the result */
					words = append(words[:index-1], words[index+1:]...)
				}
				//restarting from the beginning of the expression by breaking operators for
				break
			}
		}
	}
	num, _ = strconv.ParseFloat(words[0], 64)
	return num
}
