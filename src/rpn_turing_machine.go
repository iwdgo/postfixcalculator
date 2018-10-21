package RPN

import (
	"math"
	"strconv"
	"strings"
)

/*

Returns the result of a string in reverse polish notation (postfix) by using a turing machine.
The original band in the words exploded in a slice and results are hold on the band but in num form.
The turing band has two copies one in string and one in float. This is mandatory to avoid
multiple conversions.

A treated operation or value is erased using ? which is reserved.
A calculated value is marked as num which is reserved.

An invalid sign is interpreted as a value and the next operation panics.

*/
func RPN_Turing_machine(RPNInput string) float64 {
	words := strings.Fields(RPNInput)
	nops := len(words) //number of elements
	numbers := make([]float64, nops)
	i := 0
	leftOp, rightOp := 0.0, 0.0
	for index, w := range words {
		if strings.Contains(operatorsList, w) { // "?" is always skipped
			// at least one operand exists
			// no number can be used for detection so "num" is checked
			// if word is ?, no number there, move before
			// if word is num, already converted, read numbers
			// otherwise attempt conversion
			i = index - 1
			for words[i] != "num" && i >= 0 {
				if words[i] == "?" {
					i--
				} else if words[i] != "num" {
					if numbers[i], err = strconv.ParseFloat(words[i], 64); err != nil {
						panic("Invalid right operand")
					}
					words[i] = "num" //conversion done
				}
			}
			// w is an operator and empty cells are skipped
			if w == "sqrt" {
				// unary operator
				numbers[i] = math.Sqrt(numbers[i])
			} else {
				//binary operator
				rightOp = numbers[i]
				words[i] = "?" //erasing value
				//looking for the left operand which is to the left...
				i--
				for words[i] != "num" && i >= 0 {
					if words[i] == "?" {
						i--
					} else if words[i] != "num" {
						numbers[i], err = strconv.ParseFloat(words[i], 64)
						if err != nil {
							panic("Invalid left operand")
						}
						words[i] = "num" //conversion done
					}
				}
				leftOp = numbers[i] //cell not erased as it keeps result
				switch w {
				case "+":
					{
						numbers[i] = leftOp + rightOp
					}
				case "-":
					{
						numbers[i] = leftOp - rightOp
					}
				case "*":
					{
						numbers[i] = leftOp * rightOp
					}
				case "/":
					{
						numbers[i] = leftOp / rightOp
					}
				case "^":
					{
						numbers[i] = math.Pow(leftOp, rightOp)
					}
				default:
					panic("Invalid operator : " + w)
				}
			}
			// restarting from the beginning of the expression by breaking operators for
			words[index] = "?" // erasing operator as operation is completed
			// break was here
			index = 0 // re-starting without re-init for. No the cleanest
		}
	}
	// counting remaining ops is not needed as the previous for loop stops only
	// when words is fully explored and no operator is found
	return numbers[0] // postfix final result can only end up there
}
