// Package turingmachine implements postfix calculator using a faster Turing machine.
package turingmachine

import (
	values "github.com/iwdgo/postfixcalculator/common"
	"math"
	"strconv"
	"strings"
)

// RPNTuringMachine returns the result of a string in reverse polish notation (postfix) by using a turing machine.
// The original band in the words exploded in a slice and results are hold on the band but in num form.
// The turing band has two copies one in string and one in float. This is mandatory to avoid
// multiple conversions.
//
// A processed operation or value is erased using ? which is reserved.
// A calculated value is marked as num which is reserved.
//
// An invalid sign is interpreted as a value and the next operation panics.
func RPNTuringMachine(RPNInput string) float64 {
	words := strings.Fields(RPNInput)
	numbers := make([]float64, len(words))
	// Converting blindly is a mistake as it is very costly
	// ro is the value of the right operand
	i, ro := 0, 0.0
	var err error
	for index, w := range words {
		if strings.Contains(values.OperatorsList, w) {
			// "?" is always skipped
			// at least one operand exists
			// no number can be used for detection so "num" is checked
			// if word is ?, no number there, move before
			// if word is num, already converted, read numbers
			// otherwise attempt conversion
			i = index - 1
			for words[i] != "num" && i >= 0 {
				// former operator
				if words[i] == "?" {
					i--
					// } else if words[i] != "num" { // number not yet converted
				} else if numbers[i], err = strconv.ParseFloat(words[i], 64); err != nil {
					panic("Invalid right operand")
				} else {
					// Conversion done
					words[i] = "num"
				}
			}
			// w is an operator and empty cells are skipped
			if w == "sqrt" {
				// Unary operator
				numbers[i] = math.Sqrt(numbers[i])
			} else {
				// Binary operator
				// Copying value
				ro = numbers[i]
				// Erase operator in expression
				words[i] = "?"
				i--
				// You cannot range from max to min of index
				for words[i] != "num" && i >= 0 {
					if words[i] == "?" {
						i--
					} else if numbers[i], err = strconv.ParseFloat(words[i], 64); err != nil {
						panic("Invalid left operand")
					} else {
						// Conversion done
						words[i] = "num"
					}
				}
				switch w {
				case "+":
					numbers[i] += ro
				case "-":
					numbers[i] -= ro
				case "*":
					numbers[i] *= ro
				case "/":
					numbers[i] /= ro
				case "^":
					numbers[i] = math.Pow(numbers[i], ro)
				default:
					// TODO Never reached as no known operator was found
					panic("Invalid operator : " + w)
				}
			}
			// Complete erasure
			words[index] = "?"
			// Re-start for loop without break, nor while style. ()
			index = 0
		}
	}
	return numbers[0]
}
