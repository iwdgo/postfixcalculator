package RPN

import (
	"math"
	"strconv"
	"strings"
)

/*

Returns the result of a string in reverse polish notation (postfix) by exploding the string
to a slice and editing the slice by replacing each op by its result until only a number is left or
failing if the expression is invalid.

*/

func returnStringResult(leftOp, rightOp float64, operatorU string) string {
	switch operatorU {
	case "+":
		{
			return strconv.FormatFloat(leftOp+rightOp, 'f', 13, 64)
		}
	case "-":
		{
			return strconv.FormatFloat(leftOp-rightOp, 'f', 13, 64)
		}
	case "*":
		{
			return strconv.FormatFloat(leftOp*rightOp, 'f', 13, 64)
		}
	case "/":
		{
			return strconv.FormatFloat(leftOp/rightOp, 'f', 13, 64)
		}
	case "^":
		{
			return strconv.FormatFloat(math.Pow(leftOp, rightOp), 'f', 13, 64)
		}
	default:
		panic("Invalid operator")
	}
}

/* */
func RPN(RPNInput string) float64 {
	words := strings.Fields(RPNInput)
	index := 0
	// Known vars are needed to hold values of strconv
	num := 0.0
	num2 := 0.0
	for len(words) != 1 {
		//index = findFirstOperator(words)
		//Skipping subroutine
		err = nil
		index = 0
		for err == nil && index < len(words) {
			_, err = strconv.ParseFloat(words[index], 64) //0 iso 64 returns an int for some reason
			index++
		}
		index -= 1
		if words[index] == "sqrt" {
			//unary operator
			if num, err = strconv.ParseFloat(words[index-1], 64); err != nil {
				panic("Invalid value for sqrt")
			}
			// num = math.Sqrt(num)
			words[index-1] = strconv.FormatFloat(math.Sqrt(num), 'f', 10, 64)
			words = append(words[:index], words[index+1:]...) //removing sqrt
		} else {
			//binary operator
			if num, err = strconv.ParseFloat(words[index-2], 64); err != nil {
				panic("Invalid left operand")
			}
			if num2, err = strconv.ParseFloat(words[index-1], 64); err != nil {
				panic("Invalid right operand")
			}
			words[index-2] = returnStringResult(num, num2, words[index])
			/* removing binary operator by removing the slice until the item after the result */
			words = append(words[:index-1], words[index+1:]...)
		}
	}
	num, _ = strconv.ParseFloat(words[0], 64)
	return num
}
