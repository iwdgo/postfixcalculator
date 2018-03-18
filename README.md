**Reverse Polish Notation calculator**

Usual problem can be solved in Go using :
- slice editing (RPN)
- emulating stack using the slice (RPN emulating stack)
- using a stack-like struct (RPN stack)

All elements are providing the same features :
	- float numbers
	- arithmetic operators
	- binary operator ^
	- unary operator sqrt
	- panics on invalid values and operators

You can run examples and tests using "go test"

To run benchmark use "go test bench=."

Remarks :
- run test from inside Android Studio

