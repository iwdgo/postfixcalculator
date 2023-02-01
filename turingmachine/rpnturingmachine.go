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
	i, index, ro, lo := 0, 0, 0.0, 0.0
	var w string
	searchOperand := func() {
		for i > 0 && words[i] == "?" {
			i--
		}
	}
	panicParse := func(s string) {
		fmt.Printf("%v\n%v\n", words, numbers)
		panic(fmt.Sprintf("Invalid %s operand: %s", s, words[i]))
	}
	operationComplete := func() {
		// Mark operation as complete
		words[index] = "?"
		// Re-start for loop without break, nor while style. ()
		index = 0
	}
	var err error
	for index, w = range words {
		// Move on the band until an operator is found
		switch w {
		case "?", "num": // Ignore explicitly reserved words
		case "sqrt":
			// Unary operator
			i = index - 1
			searchOperand()
			if words[i] != "num" {
				if numbers[i], err = strconv.ParseFloat(words[i], 64); err != nil {
					panicParse("unique")
				}
				words[i] = "num"
			}
			numbers[i] = math.Sqrt(numbers[i])
			operationComplete()
		case "+", "-", "*", "/", "^":
			i = index - 1
			searchOperand()
			// Load ro with right operand
			if words[i] == "num" {
				ro = numbers[i]
			} else {
				if ro, err = strconv.ParseFloat(words[i], 64); err != nil {
					panicParse("right")
				}
			}
			// Binary operator
			words[i] = "?"
			i--
			searchOperand()
			if words[i] == "num" {
				lo = numbers[i]
			} else {
				if lo, err = strconv.ParseFloat(words[i], 64); err != nil {
					panicParse("left")
				}
				// Mark operand as converted in band. It will hold the result of the operator.
				words[i] = "num"
			}
			switch w {
			case "+":
				numbers[i] = lo + ro
			case "-":
				numbers[i] = lo - ro
			case "*":
				numbers[i] = lo * ro
			case "/":
				numbers[i] = lo / ro
			case "^":
				numbers[i] = math.Pow(lo, ro)
			default:
				// TODO Never reached as no unknown operator can arrive here
				panic(fmt.Sprintf("Invalid operator: %s", w))
			}
			operationComplete()
		}
	}
	return numbers[0]
}
