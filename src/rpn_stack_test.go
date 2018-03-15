package RPN

import (
	"fmt"
	"testing"
)

/*
Running a single test
*/
func ExampleRPN_stack() {
	fmt.Printf("%.13f", RPN_stack(input))
	// Output: 3.0001220703125
}

/* Debug output when uncommenting fmt is
For postfix "3 4 2 * 1 5 - 2 3 ^ ^ / +"

Token            Action            Stack
  3    Push num onto top of stack  3
  4    Push num onto top of stack  4
  2    Push num onto top of stack  2
  *    Apply op to top of stack    8
  1    Push num onto top of stack  1
  5    Push num onto top of stack  5
  -    Apply op to top of stack    -4
  2    Push num onto top of stack  2
  3    Push num onto top of stack  3
  ^    Apply op to top of stack    8
  ^    Apply op to top of stack    65536
  /    Apply op to top of stack    0.0001220703125
  +    Apply op to top of stack    3.0001220703125

The final value is 3.0001220703125
*/

func TestRPN_stack(t *testing.T) {
	actual := RPN_stack(input)
	if actual != 3.0001220703125 {
		t.Errorf("RPN(%s): expected %f, actual %f", input, 3.0001220703125, actual)
	}
}

/*
go test -bench=. allows to test the func Test1
*/
func BenchmarkRPN_stack(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		RPN_stack(input)
	}
}
