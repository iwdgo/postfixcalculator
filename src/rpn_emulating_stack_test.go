package rpn

import (
	"fmt"
	"testing"
)

/*
Running a single test
*/
func ExampleRPNEmulatingStack() {
	fmt.Printf("%.13f", RPNEmulatingStack(input))
	// Output: 3.0001220703125
}

func ExampleRPNEmulatingStack_sqrt() {
	fmt.Printf("%f", RPNEmulatingStack(RPNInput))
	// Output: 30.000000
}

func TestRPNEmulatingStack(t *testing.T) {
	if got := RPNEmulatingStack(input); got != InputWant {
		t.Errorf("RPN(%s): got %f, want %f", input, got, InputWant)
	}
}

// To check panic if expression is invalid.
func TestRPNEmulatingStack_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	RPNEmulatingStack(InvalidInput)
}

func BenchmarkRPNEmulatingStack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNEmulatingStack(input)
	}
}
