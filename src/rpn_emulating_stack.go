package RPN

import (
	"math"
	"strconv"
	"strings"
	//"fmt"
)

/* Prints the result of a string in reverse polish notation (postfix) using stack emulation
This emulation is incorrect as the stack is accessed sequentially when only pop/push are allowed
http://rosettacode.org/wiki/Parsing/RPN_calculator_algorithm#Go

Supporting
	- float
	- arithmetic operators
	- binary operator ^
	- panics on invalid values and operators

Extensions :
	- no unary operator
*/
//const RPNInput = "6 8 4.0 sqrt + 3.01 1.99 + - *" //= 30
//const input = "3 4 2 * 1 5 - 2 3 ^ ^ / +" //= 3.0001220703125
// 6 10 5 - *
// 6 5 * = 42
//const RPNInput = "6 8 +"
//var err error

func RPN_emulating_stack(input string) float64 {
	//fmt.Printf("For postfix %q\n", input)
	//fmt.Println("\nToken            Action            Stack")
	var stack []float64 //Only contains numbers
	for _, tok := range strings.Fields(input) {
		//action := "Apply op to top of stack"
		switch tok {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "^":
			stack[len(stack)-2] =
				math.Pow(stack[len(stack)-2], stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		default:
			//action = "Push num onto top of stack"
			f, _ := strconv.ParseFloat(tok, 64)
			//If value is not a number, no error caught and 0 is pushed on stack
			stack = append(stack, f)
		}
		//fmt.Printf("%3s    %-26s  %v\n", tok, action, stack)
	}
	//fmt.Println("\nThe final value is", stack[0])
	return stack[0]
}
