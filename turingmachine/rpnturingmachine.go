// Package turingmachine implements postfix calculator using a faster Turing machine.
package turingmachine

import (
	"fmt"
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
	// A slice of words contains the expression.
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
		for words[i] != "num" {
			// if word is ?, no number there, move before
			if words[i] == "?" {
				i--
				continue
			}
			if numbers[i], err = strconv.ParseFloat(words[i], 64); err != nil {
				fmt.Printf("%v\n", words)
				fmt.Printf("%v\n", numbers)
				panic(fmt.Sprintf("Invalid right operand: %s", w))
			}
			// Mark operand as converted in band
			words[i] = "num"
		}
		// w is an operator and empty cells are skipped
		if w == "sqrt" {
			// Unary operator
			numbers[i] = math.Sqrt(numbers[i])
		} else {
			// Binary operator
			// Copying value
			ro = numbers[i]
			// Number is consumed
			words[i] = "?"
			i--
			for words[i] != "num" {
				if words[i] == "?" {
					i--
					continue
				}
				if numbers[i], err = strconv.ParseFloat(words[i], 64); err != nil {
					fmt.Printf("%v\n", words)
					fmt.Printf("%v\n", numbers)
					panic(fmt.Sprintf("Invalid left operand: %s", w))
				}
				// Mark operand as converted in band. It will hold the result of the operator.
				words[i] = "num"
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
				panic(fmt.Sprintf("Invalid operator: %s", w))
			}
		}
		// Mark operation as complete
		words[index] = "?"
		// Re-start for loop without break, nor while style. ()
		index = 0
	}
	return numbers[0]
}
