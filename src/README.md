**Reverse Polish Notation calculator**

Usual problem can be solved in Go using :
- slice editing (RPN)
- emulating stack using the slice (RPN emulating stack)
- using a stack-like struct (RPN stack)

You can run examples and tests using "go test"
Verbose print has been removed print has been removed.

To run benchmark use "go test bench=."
c:\Users\Costa\Documents\Google\CloudSDK\sandbox\samples\RPN>go test -bench=.
goos: windows
goarch: amd64
BenchmarkRPN_emulating_stack-4           1000000              1131 ns/op
BenchmarkRPN_stack-4                     1000000              1705 ns/op
BenchmarkRPN-4                            500000              3696 ns/op
PASS
ok      _/c_/Users/Costa/Documents/Google/CloudSDK/sandbox/samples/RPN  5.295s

Remarks :
- align calculator capabilities as RPN handles unary operator sqrt
- run test from inside Android Studio

