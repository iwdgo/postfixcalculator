// Package turingmachine implements postfix calculator using a faster Turing machine.
package turingmachine

import (
	values "github.com/iwdgo/postfixcalculator/common"
	"math"
	"strconv"
	"strings"
)

// RPNTuringMachine returns the result of a string in reverse polish notation (postfix) using a turing machine.
// It expects a band as a string read from left to write complying with postfix notation where blank
// spaces operands and operators.
// A first slice is the band where words are held in a slice. A processed operator is erased by replacing
// it with ? in the expression band. Similarly, a converter number is erased using 'num'.
// Both ? and num are reserved keywords.
// A second slice holds numbers as float using the same index as the band. It avois multiple
// conversions and searches.
// An invalid sign is interpreted as a value and the next operation panics.
func RPNTuringMachine(RPNInput string) float64 {
	words := strings.Fields(RPNInput)
	numbers := make([]float64, len(words))
	// ro is the value of the right operand
	i, ro := 0, 0.0
	var err error
	for index, w := range words {
		if strings.Contains(values.OperatorsList, w) {
			// At least one operand is expected in a preceding place.
			i = index - 1
			for words[i] != "num" && i >= 0 {
				// if word is ?, no number there, move before
				if words[i] == "?" {
					i--
					// if word is num, already converted, read numbers
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
