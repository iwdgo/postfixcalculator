// Package turingmachine implements postfix calculator using a faster Turing machine.
package turingmachine

import (
	"fmt"
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
	// A slice of words contains the expression.
	words := strings.Fields(RPNInput)
	// A slice of floats contains the values to avoid repeating conversion and words.
	numbers := make([]float64, len(words))
	// i is the index in the current operation
	// lo and ro hold values of left and right operand as writing to the slice is more expensive
	i, index, lastIndex, ro, lo := 0, 0, len(words)-1, 0.0, 0.0
	var err error
	for index = range words {
		// Move on the band until an operator is found
		switch words[index] {
		case "?": // Ignore explicitly reserved words
		case "sqrt":
			// Unary operator
			i = index - 1
			for i > 0 && words[i] == "?" {
				i--
			}
			numbers[i] = math.Sqrt(numbers[i])
			if index == lastIndex {
				return numbers[0]
			}
			words[index] = "?"
			index = 0
		case "+":
			i = index - 1
			for i > 0 && words[i] == "?" {
				i--
			}
			ro = numbers[i]
			words[i] = "?"
			for i > 0 && words[i] == "?" {
				i--
			}
			lo = numbers[i]
			numbers[i] = lo + ro
			if index == lastIndex {
				return numbers[0]
			}
			words[index] = "?"
			index = 0
		case "-":
			i = index - 1
			for i > 0 && words[i] == "?" {
				i--
			}
			ro = numbers[i]
			words[i] = "?"
			for i > 0 && words[i] == "?" {
				i--
			}
			lo = numbers[i]
			numbers[i] = lo - ro
			if index == lastIndex {
				return numbers[0]
			}
			words[index] = "?"
			index = 0
		case "*":
			i = index - 1
			for i > 0 && words[i] == "?" {
				i--
			}
			ro = numbers[i]
			words[i] = "?"
			for i > 0 && words[i] == "?" {
				i--
			}
			lo = numbers[i]
			numbers[i] = lo * ro
			if index == lastIndex {
				return numbers[0]
			}
			words[index] = "?"
			index = 0
		case "/":
			i = index - 1
			for i > 0 && words[i] == "?" {
				i--
			}
			ro = numbers[i]
			words[i] = "?"
			for i > 0 && words[i] == "?" {
				i--
			}
			lo = numbers[i]
			numbers[i] = lo / ro
			if index == lastIndex {
				return numbers[0]
			}
			words[index] = "?"
			index = 0
		case "^":
			// Binary operator
			i = index - 1
			for i > 0 && words[i] == "?" {
				i--
			}
			ro = numbers[i]
			words[i] = "?"
			for i > 0 && words[i] == "?" {
				i--
			}
			lo = numbers[i]
			numbers[i] = math.Pow(lo, ro)
			if index == lastIndex {
				return numbers[0]
			}
			words[index] = "?"
			index = 0
		default:
			// Not a known operator, it must be an operand
			if numbers[index], err = strconv.ParseFloat(words[index], 64); err != nil {
				fmt.Printf("%v\n%v\n", words, numbers)
				panic(fmt.Sprintf("Invalid operator or operand: %s", words[i]))
			}
		}
	}
	return numbers[0]
}
