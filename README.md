**Reverse Polish Notation calculator**

Usual problem can be solved in Go using :
- slice editing (RPN)
- emulating stack using the slice (RPN emulating stack)
- using a stack-like struct (RPN stack)
- using an operators list as a string (RPN operators list)
- using the same operators list and the string package (RPN operators no fields)

All elements are providing the same features :
	- float numbers
	- arithmetic operators
	- binary operator ^
	- unary operator sqrt
	- panics on invalid values and operators

You can run examples and tests using "go test"

To run benchmark use "go test bench=."

Emulating the stack is the most efficient method.
Using a struct in a provided Go Doc package is already costly.
As soon as you add slice editing, you divide performance by 3.

goos: windows
goarch: amd64
BenchmarkRPN_emulating_stack-4 | 2000000 | 961 ns/op
BenchmarkRPN_operators_list-4 | 500000 | 3209 ns/op
BenchmarkRPN_operators_no_fields-4 | 500000 | 2892 ns/op
BenchmarkRPN_stack-4 | 1000000 | 1884 ns/op
BenchmarkRPN-4 | 500000 | 3988 ns/op
PASS

