/* RPNOperatorsList returns the result of a string in reverse polish notation (postfix) by exploding the string
to a slice and editing the slice by replacing each op by its result until only a number is left or
failing if the expression is invalid.

The expression is kept in string format. Multiple conversions to test the validity of the expression
is lowering efficiency. A list of authorized operators is kept to find the fist one in the list.

Removing subroutines has no visible impact here.

An invalid sign is interpreted as a value and the next operation panics.
*/
package operatorslist

import (
	values "github.com/iwdgo/postfixcalculator/common"
	"math"
	"strconv"
	"strings"
)

func RPNOperatorsList(RPNInput string) float64 {
	operators := strings.Fields(values.OperatorsList)
	words := strings.Fields(RPNInput)
	num, leftOp, rightOp := 0.0, 0.0, 0.0
	var err error
	operationCompleted := false
	for len(words) != 1 { // length of expression stops processing
		err = nil
		operationCompleted = false
		for index, w := range words {
			for _, op := range operators {
				if w == op { // if an operator is found, calculate value otherwise keep searching.
					if w == "sqrt" {
						// unary operator
						if num, err = strconv.ParseFloat(words[index-1], 64); err != nil {
							panic("Invalid value for sqrt")
						}
						// num = math.Sqrt(num)
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
						switch op {
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

						// removing binary operator by removing the slice until the item after the result.
						words = append(words[:index-1], words[index+1:]...)
					}
					// restarting from the beginning of the expression by breaking both loops
					operationCompleted = true
					break
				}
			}
			if operationCompleted {
				break
			}
		}
	}
	num, _ = strconv.ParseFloat(words[0], 64)
	return num
}
