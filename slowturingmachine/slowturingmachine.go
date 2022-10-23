/*
	Package slowturingmachine implements postfix calculator using a Turing machine using a for loop which is

rewinded without initialization.

Code is shorter that the fast version of the Turing machine as it is using a "while" format.
It remains slower than the usual for structure.
*/
package slowturingmachine

import (
	values "github.com/iwdgo/postfixcalculator/common"
	"math"
	"strconv"
	"strings"
)

// RPNSlowTuringMachine returns the result of a string in reverse polish notation (postfix)
// using a Turing machine.
//
// The original band in the words exploded in a slice and results are held on the band but in num form.
// The band has two copies one in string and one in float. This is mandatory to avoid costly conversions.
// A processed operation or value is erased by replacing it with ? which is a reserved sign.
// A calculated value is marked as num which is also a reserved word.
// An invalid sign is interpreted as a value and the next operation panics.
func RPNSlowTuringMachine(RPNInput string) float64 {
	words := strings.Fields(RPNInput)
	numbers := make([]float64, len(words))
	i, index := 0, 0
	rightOp := 0.0
	var err error
	for index < len(words) {
		w := words[index]
		if strings.Contains(values.OperatorsList, w) { //"?" is always skipped
			//at least one operand exists
			//no number can be used for detection so "num" is checked
			//if word is ?, no number there, move before
			//if word is num, already converted, read numbers
			//otherwise attempt conversion
			i = index - 1
			for words[i] != "num" && i >= 0 {
				if words[i] == "?" {
					i--
				} else if words[i] != "num" {
					numbers[i], err = strconv.ParseFloat(words[i], 64)
					if err != nil {
						panic("Invalid right operand")
					}
					// Conversion done
					words[i] = "num"
				}
			}
			// w is an operator
			if w == "sqrt" {
				// Unary operator
				numbers[i] = math.Sqrt(numbers[i])
			} else {
				// Binary operator
				rightOp = numbers[i]
				// Erase value
				words[i] = "?"
				// Looking for the left operand which is to the left...
				i--
				for words[i] != "num" && i >= 0 {
					if words[i] == "?" {
						i--
					} else if words[i] != "num" {
						numbers[i], err = strconv.ParseFloat(words[i], 64)
						if err != nil {
							panic("Invalid left operand")
						}
						// Conversion done
						words[i] = "num"
					}
				}
				// Cell will store result
				switch w {
				case "+":
					numbers[i] += rightOp
				case "-":
					numbers[i] -= rightOp
				case "*":
					numbers[i] *= rightOp
				case "/":
					numbers[i] /= rightOp
				case "^":
					numbers[i] = math.Pow(numbers[i], rightOp)
				default:
					panic("Invalid operator : " + w)
				}
			}
			// Restarting from the beginning of the expression by breaking operators for
			// Erase operator
			words[index] = "?"
			// Re-start without re-init for loop
			index = 0
		} else {
			index++
		}

	}
	return numbers[0]
}
