package common

import (
	"os"
	"testing"
)

const OperatorsList = "+ - * / ^ sqrt"

// Two examples and their expected results
const (
	RPNInput     = "6 8 4.0 sqrt + 3.01 1.99 + - *" //= 30.000000
	Input        = "3 4 2 * 1 5 - 2 3 ^ ^ / +"      //= 3.0001220703125
	InvalidInput = "3 4 2 * ? 1 5 - 2 3 ^ ^ / +"    // No value

	RPNInputWant = 30.000000
	InputWant    = 3.0001220703125
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}