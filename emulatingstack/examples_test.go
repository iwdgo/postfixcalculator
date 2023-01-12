package emulatingstack_test

import (
	"fmt"
	"github.com/iwdgo/postfixcalculator/emulatingstack"
)

func ExampleRPNEmulatingStack() {
	fmt.Printf("%f", emulatingstack.RPNEmulatingStack("6 8 4.0 sqrt + 3.01 1.99 + - *"))
	// Output: 30.000000
}

func ExampleRPNEmulatingStack_exp() {
	fmt.Printf("%.13f", emulatingstack.RPNEmulatingStack("3 4 2 * 1 5 - 2 3 ^ ^ / +"))
	// Output: 3.0001220703125
}
