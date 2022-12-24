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
// Both ? and num are reserved keywords.
// Method will panic on failed conversions or unknown operators.
func RPNTuringMachine(RPNInput string) float64 {
	// A slice of words contains the expression. A processed operator is erased by replacing
	// it with ? in the expression. Similarly, a converted number is erased using 'num'.
	words := strings.Fields(RPNInput)
	// A slice of floats contains the values to avoid repeating conversion and words.
	numbers := make([]float64, len(words))
	// i is the index in the current operation
	// ro is the value of the right operand
	i, ro := 0, 0.0
	var err error
	for index, w := range words {
		if !strings.Contains(values.OperatorsList, w) {
			continue
		}
		// At least one operand is expected in a preceding place.
		i = index - 1
		for {
			// if word is ?, no number there, move before
			if words[i] == "?" {
				i--
				continue
			}
			if words[i] == "num" {
				break
			}
			if numbers[i], err = strconv.ParseFloat(words[i], 64); err != nil {
				panic("Invalid right operand")
			}
			// Conversion done
			words[i] = "num"
			break
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
			for {
				// if word is ?, no number there, move before
				if words[i] == "?" {
					i--
					continue
				}
				if words[i] == "num" {
					break
				}
				if numbers[i], err = strconv.ParseFloat(words[i], 64); err != nil {
					panic("Invalid left operand")
				}
				// Conversion done
				words[i] = "num"
				break
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
	return numbers[0]
}
