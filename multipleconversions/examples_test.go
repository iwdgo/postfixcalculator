package multipleconversions_test

import (
	"fmt"
	"github.com/iwdgo/postfixcalculator/multipleconversions"
)

func ExampleRPN() {
	fmt.Printf("%f", multipleconversions.RPN("6 8 4.0 sqrt + 3.01 1.99 + - *"))
	// Output: 30.000000
}

func ExampleRPN_exp() {
	fmt.Printf("%.13f", multipleconversions.RPN("3 4 2 * 1 5 - 2 3 ^ ^ / +"))
	// Output: 3.0001220703125
}
