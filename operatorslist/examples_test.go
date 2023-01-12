package operatorslist_test

import (
	"fmt"
	"github.com/iwdgo/postfixcalculator/operatorslist"
)

func ExampleRPNOperatorsList() {
	fmt.Printf("%f", operatorslist.RPNOperatorsList("6 8 4.0 sqrt + 3.01 1.99 + - *"))
	// Output: 30.000000
}

func ExampleRPNOperatorsList_exp() {
	fmt.Printf("%.13f", operatorslist.RPNOperatorsList("3 4 2 * 1 5 - 2 3 ^ ^ / +"))
	// Output: 3.0001220703125
}
