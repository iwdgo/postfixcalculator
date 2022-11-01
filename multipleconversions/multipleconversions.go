// Package multipleconversions implements postfix calculator using conversions.
package multipleconversions

import (
	"math"
	"strconv"
	"strings"
)

// RPN returns the result of an expression using reverse polish notation (postfix) by exploding the string
// to a slice and editing the slice by replacing each op by its result until only a number is left or
// failing if the expression is invalid.
func RPN(RPNInput string) float64 {
	words := strings.Fields(RPNInput)
	index := 0
	// Known vars are needed to hold values of strconv
	var num, num2 float64
	var err error
	for len(words) != 1 {
		err = nil
		index = 0
		// Search for an operator by detecting a failed conversion.
		// An invalid operand also triggers a calculation, i.e. no conversion can fail during the calculation
		for err == nil && index < len(words) {
			_, err = strconv.ParseFloat(words[index], 64)
			if err == nil {
				index++
			}
		}
		if words[index] == "sqrt" {
			// Unary operator for which conversion of operand cannot fail
			num, _ = strconv.ParseFloat(words[index-1], 64)
			// Result is overwriting the initial value and stored as a string
			words[index-1] = strconv.FormatFloat(math.Sqrt(num), 'f', 10, 64)
			// Remove operator sqrt
			words = append(words[:index], words[index+1:]...)
		} else {
			// Binary operator
			num, _ = strconv.ParseFloat(words[index-2], 64)
			num2, _ = strconv.ParseFloat(words[index-1], 64)
			switch words[index] {
			case "+":
				words[index-2] = strconv.FormatFloat(num+num2, 'f', 13, 64)
			case "-":
				words[index-2] = strconv.FormatFloat(num-num2, 'f', 13, 64)
			case "*":
				words[index-2] = strconv.FormatFloat(num*num2, 'f', 13, 64)
			case "/":
				words[index-2] = strconv.FormatFloat(num/num2, 'f', 13, 64)
			case "^":
				words[index-2] = strconv.FormatFloat(math.Pow(num, num2), 'f', 13, 64)
			default:
				panic("Invalid operator")
			}
			// Remove operator and right operand
			words = append(words[:index-1], words[index+1:]...)
		}
	}
	num, err = strconv.ParseFloat(words[0], 64)
	if err != nil {
		panic("Invalid operand or expression with one operator")
	}
	return num
}
