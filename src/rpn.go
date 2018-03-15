package RPN

import (
	"math"
	"strconv"
	"strings"
	//"fmt"
)

/* Prints the result of a string in reverse polish notation (postfix) by exploding the string
to a slice and editing the slice by replacing each op by its result until only a number is left

Supporting
	- float
	- arithmetic operators
	- unary operator sqrt
	- panics on invalid values and operators

Possible additions :
	- exponent ^
	- input value (read entry)
	-
*/
//const RPNInput = "6 8 4.0 sqrt + 3.01 1.99 + - *" //= 30
// 6 10 5 - *
// 6 5 * = 42
//const RPNInput = "6 8 +"
//var err error

func findFirstOperator(operands []string) int {
	err = nil
	i := 0
	for err == nil && i < len(operands) {
		_, err = strconv.ParseFloat(operands[i], 64) //0 returns an int for some reason
		i++
	}
	return i - 1
}

func returnStringResult(leftOp, rightOp float64, operatorU string) string {
	switch operatorU {
	case "+":
		{
			return strconv.FormatFloat(leftOp+rightOp, 'f', 10, 64)
		}
	case "-":
		{
			return strconv.FormatFloat(leftOp-rightOp, 'f', 10, 64)
		}
	case "*":
		{
			return strconv.FormatFloat(leftOp*rightOp, 'f', 10, 64)
		}
	case "/":
		{
			return strconv.FormatFloat(leftOp/rightOp, 'f', 10, 64)
		}
	default:
		panic("Invalid operator")
	}
}

/* */
func RPN(RPNInput string) float64 {
	words := strings.Fields(RPNInput)
	index := 0
	num := 0.0
	num2 := 0.0
	for len(words) != 1 {
		index = findFirstOperator(words)
		if words[index] == "sqrt" {
			//unary operator
			num, err = strconv.ParseFloat(words[index-1], 64)
			if err != nil {
				panic("Invalid value for sqrt")
			}
			num = math.Sqrt(num)
			words[index-1] = strconv.FormatFloat(num, 'f', 10, 64)
			words = append(words[:index], words[index+1:]...) //removing sqrt
		} else {
			//binary operator
			num, err = strconv.ParseFloat(words[index-2], 64)
			if err != nil {
				panic("Invalid left operand")
			}
			num2, err = strconv.ParseFloat(words[index-1], 64)
			if err != nil {
				panic("Invalid right operand")
			}
			words[index-2] = returnStringResult(num, num2, words[index])
			/* removing binary operator by removing the slice until the item after the result */
			words = append(words[:index-1], words[index+1:]...)
		}
	}
	//fmt.Println(words[0]) //fmt is needed to use the test function
	num, _ = strconv.ParseFloat(words[0], 64)
	return num
}
