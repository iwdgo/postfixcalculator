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
		for err == nil && index < len(words) {
			_, err = strconv.ParseFloat(words[index], 64)
			index++
		}
		// TODO err is ignored expecting an operator but could be an operand
		index -= 1
		if words[index] == "sqrt" {
			// Unary operator
			if num, err = strconv.ParseFloat(words[index-1], 64); err != nil {
				// TODO Unreachable
				panic("Invalid value for sqrt")
			}
			words[index-1] = strconv.FormatFloat(math.Sqrt(num), 'f', 10, 64)
			// Remove operator sqrt
			words = append(words[:index], words[index+1:]...)
		} else {
			// Binary operator
			if num, err = strconv.ParseFloat(words[index-2], 64); err != nil {
				panic("Invalid left operand")
			}
			if num2, err = strconv.ParseFloat(words[index-1], 64); err != nil {
				panic("Invalid right operand")
			}
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
				// TODO Unreachable
				panic("Invalid operator")
			}
			// Remove operator and right operand
			words = append(words[:index-1], words[index+1:]...)
		}
	}
	num, err = strconv.ParseFloat(words[0], 64)
	if err != nil {
		panic("Invalid operator or operand")
	}
	return num
}
