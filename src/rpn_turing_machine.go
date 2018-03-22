package RPN

import (
	"math"
	"strconv"
	"strings"
)

/*

Returns the result of a string in reverse polish notation (postfix) by using a turing machine.
The original band in the words exploded in a slice and results are hold on the band but in num form.
The turing band has two copies one in string and one in float. This is mandatory to avoid
multiple conversions.

A treated operation or value is erased using ? which is reserved.
A calculated value is marked as num which is reserved.

An invalid sign is interpreted as a value and the next operation panics.

*/
/* */
func RPN_Turing_machine(RPNInput string) float64 {
	words := strings.Fields(RPNInput)
	nops := len(words) //number of elements
	numbers := make([]float64,len(words))
	i := 0
	num := 0.0
	nop := 0
	leftOp := 0.0
	rightOp := 0.0
	for nop != nops - 1 { //length of expression stops processing
		err = nil
		for index, w := range words {
			if strings.Contains(operatorsList, w) {//"?" is always skipped
				//w is an operator and empty cells are skipped
				if w == "sqrt" {
					//unary operator
					i = index - 1
					num = 0.0
					for num == 0.0 && i >= 0 {
						//if word is ?, no number there, move before
						//if word is num, already converted, read numbers
						//otherwise attempt conversion
						if (words[i] == "?") {
							i--
						} else if words[i] == "num" {
							num = numbers[i] //will be overridden
						} else {
							num, err = strconv.ParseFloat(words[i], 64)
							if err != nil {
								panic("Invalid sqrt operand")
							}
							words[i] = "num"
						}
					}
					numbers[i] = math.Sqrt(num)
				} else {
					//binary operator
					i = index - 1
					rightOp = 0.0
					for rightOp == 0.0 && i >= 0  {
						if (words[i] == "?") {
							i--
						} else if words[i] == "num" {
							rightOp = numbers[i]
							words[i] = "?" //erasing value
						} else {
							rightOp, err = strconv.ParseFloat(words[i], 64)
							if err != nil {
								panic("Invalid right operand")
							}
							//numbers[index-1] =
							words[i] = "?" //erasing value
						}
					}
					//looking for the left operand which is to the left
					i--
					leftOp = 0.0
					for leftOp == 0.0 && i >= 0  {
						if (words[i] == "?") {
							i--
						} else if words[i] == "num" {
							leftOp = numbers[i]
						} else {
							leftOp, err = strconv.ParseFloat(words[i], 64)
							if err != nil {
								panic("Invalid right operand")
							}
							words[i] = "num" //result will be stored
						}
					}

					switch w {
					case "+":
						{
							numbers[i] = leftOp+rightOp
						}
					case "-":
						{
							numbers[i] = leftOp-rightOp
						}
					case "*":
						{
							numbers[i] = leftOp*rightOp
						}
					case "/":
						{
							numbers[i] = leftOp/rightOp
						}
					case "^":
						{
							numbers[i] = math.Pow(leftOp, rightOp)
						}
					default:
						panic("Invalid operator : "+w)
					}
				}
				//restarting from the beginning of the expression by breaking operators for
				words[index] = "?" //erasing operator as operation is completed
				break
			}
		}
		//counting remaining ops.
		nop = 0 //very costly strings.Count(strings.Join(words," "),"?")
		for _, w := range words {
			if w == "?" {
				nop += 1
			}
		}//nop = nops - 1 to exit, i.e. [num ?...?]

	}
	return numbers[0] //postfix final result can only end up there
}
