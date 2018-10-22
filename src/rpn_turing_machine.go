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

A processed operation or value is erased using ? which is reserved.
A calculated value is marked as num which is reserved.

An invalid sign is interpreted as a value and the next operation panics.

*/
func RPN_Turing_machine(RPNInput string) float64 {
	words := strings.Fields(RPNInput)
	numbers := make([]float64, len(words))
	// Convertinc blindly is a mistake as it is very costly
	i, ro := 0, 0 // ro is the index of the right operand
	for index, w := range words {
		if strings.Contains(operatorsList, w) { // "?" is always skipped
			// at least one operand exists
			// no number can be used for detection so "num" is checked
			// if word is ?, no number there, move before
			// if word is num, already converted, read numbers
			// otherwise attempt conversion
			i = index - 1
			for words[i] != "num" && i >= 0 {
				if words[i] == "?" { // former operator
					i--
				} else if words[i] != "num" { // number not yet converted
					if numbers[i], err = strconv.ParseFloat(words[i], 64); err != nil {
						panic("Invalid right operand")
					}
					words[i] = "num" // conversion done
				} // else operand is found
			}
			// w is an operator and empty cells are skipped
			if w == "sqrt" {
				// unary operator
				numbers[i] = math.Sqrt(numbers[i])
			} else {
				// binary operator
				ro = i
				words[i] = "?" // erasing value in operation string
				// looking for the left operand which is to the left...
				i--
				// You can't range from max to min of index
				for words[i] != "num" && i >= 0 {
					if words[i] == "?" {
						i--
					} else if words[i] != "num" {
						if numbers[i], err = strconv.ParseFloat(words[i], 64); err != nil {
							panic("Invalid left operand")
						}
						words[i] = "num" // conversion done
					}
				}
				switch w {
				case "+":
					{
						numbers[i] += numbers[ro]
					}
				case "-":
					{
						numbers[i] -= numbers[ro]
					}
				case "*":
					{
						numbers[i] *= numbers[ro]
					}
				case "/":
					{
						numbers[i] /= numbers[ro]
					}
				case "^":
					{
						numbers[i] = math.Pow(numbers[i], numbers[ro])
					}
				default:
					panic("Invalid operator : " + w)
				}
			}
			words[index] = "?" // erasing operator as operation is completed
			// break was here
			index = 0 // re-starting without re-init for. Not the cleanest
		}
	}
	// counting remaining ops is not needed as the previous for loop stops only
	// when words is fully explored and no operator is found
	return numbers[0] // postfix final result can only end up there
}
