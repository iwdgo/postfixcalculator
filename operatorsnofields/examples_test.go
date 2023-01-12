package operatorsnofields_test

import (
	"fmt"
	"github.com/iwdgo/postfixcalculator/operatorsnofields"
)

func ExampleRPNOperatorsNoFields() {
	fmt.Printf("%f", operatorsnofields.RPNOperatorsNoFields("6 8 4.0 sqrt + 3.01 1.99 + - *"))
	// Output: 30.000000
}

func ExampleRPNOperatorsNoFields_exp() {
	fmt.Printf("%.13f", operatorsnofields.RPNOperatorsNoFields("3 4 2 * 1 5 - 2 3 ^ ^ / +"))
	// Output: 3.0001220703125
}
