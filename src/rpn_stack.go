package RPN

import (
	"math"
	"strconv"
	"strings"
	//"fmt"
	"github.com/golang-collections/collections/stack"
)

/* Prints the result of a string in reverse polish notation (postfix) using stack package from collections

Supporting
	- float
	- arithmetic operators
	- unary operator ^
	- panics on invalid values and operators

Standard output :
For postfix "3 4 2 * 1 5 - 2 3 ^ ^ / +"

Token            Action            Stack
  3    Push num onto top of stack  [3]
  4    Push num onto top of stack  [3 4]
  2    Push num onto top of stack  [3 4 2]
  *    Apply op to top of stack    [3 8]
  1    Push num onto top of stack  [3 8 1]
  5    Push num onto top of stack  [3 8 1 5]
  -    Apply op to top of stack    [3 8 -4]
  2    Push num onto top of stack  [3 8 -4 2]
  3    Push num onto top of stack  [3 8 -4 2 3]
  ^    Apply op to top of stack    [3 8 -4 8]
  ^    Apply op to top of stack    [3 8 65536]
  /    Apply op to top of stack    [3 0.0001220703125]
  +    Apply op to top of stack    [3.0001220703125]

The final value is 3.0001220703125
*/
//const RPNInput = "6 8 4.0 sqrt + 3.01 1.99 + - *" //= 30
//const input = "3 4 2 * 1 5 - 2 3 ^ ^ / +" //= 3.0001220703125
// 6 10 5 - *
// 6 5 * = 42
//const RPNInput = "6 8 +"
//var err error

func RPN_stack(input string) float64 {
	//fmt.Printf("For postfix %q\n", input)
	//fmt.Println("\nToken            Action            Stack")
	num := 0.0
	stackOperands := stack.Stack{} //= stack.New()
	for _, tok := range strings.Fields(input) {
		//action := "Apply op to top of stack"
		switch tok {
		case "+":
			stackOperands.Push(stackOperands.Pop().(float64) + stackOperands.Pop().(float64))
		case "-":
			stackOperands.Push(-stackOperands.Pop().(float64) + stackOperands.Pop().(float64))
		case "*":
			stackOperands.Push(stackOperands.Pop().(float64) * stackOperands.Pop().(float64))
		case "/":
			stackOperands.Push(1 / stackOperands.Pop().(float64) * stackOperands.Pop().(float64))
		case "^":
			num = stackOperands.Pop().(float64)
			stackOperands.Push(math.Pow(stackOperands.Pop().(float64), num))
		default:
			//action = "Push num onto top of stack"
			f, _ := strconv.ParseFloat(tok, 64)
			stackOperands.Push(f)
		}
		//fmt.Printf("%3s    %-26s  %v\n", tok, action, stackOperands.Peek().(float64))
	}
	//fmt.Println("\nThe final value is", stackOperands.Peek().(float64))
	return stackOperands.Peek().(float64) //Dijkstra would pop
}
