// RPNStack prints the result of a string in reverse polish notation (postfix) using stack package
// from collections.
package stack

import (
	"github.com/badgerodon/collections/stack"
	"math"
	"strconv"
	"strings"
)

func RPNStack(input string) float64 {
	num := 0.0
	stackOperands := stack.Stack{} //= stack.New()
	var err error
	for _, tok := range strings.Fields(input) {
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
			// stackOperands.Push(math.Pow(stackOperands.Pop().(float64), stackOperands.Pop().(float64))) fails
			// Type casting is probably not executed correctly
		case "sqrt":
			stackOperands.Push(math.Sqrt(stackOperands.Pop().(float64)))
		default:
			f, _ := strconv.ParseFloat(tok, 64)
			if err != nil {
				panic("Invalid number and not a known operator")
			}
			stackOperands.Push(f)
		}
	}
	return stackOperands.Peek().(float64) //Dijkstra would pop
}
