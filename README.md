[![GoDoc](https://godoc.org/github.com/iWdGo/postfixcalculator?status.svg)](https://godoc.org/github.com/iWdGo/postfixcalculator)
[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/postfixcalculator)](https://goreportcard.com/report/github.com/iwdgo/postfixcalculator)
[![codecov](https://codecov.io/gh/iWdGo/postfixcalculator/branch/master/graph/badge.svg)](https://codecov.io/gh/iWdGo/postfixcalculator)

# Reverse Polish Notation calculator

Several solutions to this well-known problem show dramatic differences depending on the approach.
A Turing machine emulation beats all of them.

***Features***

All solutions support :
- float numbers
- arithmetic operators
- binary operator ^ (exponent)
- unary operator sqrt

Input is a string. Process panics on invalid values and operators.

***Solutions***

- slice editing (RPN) using [stack package](https://godoc.org/github.com/golang-collections/collections/stack)
  requires to `go get github.com/golang-collections/collections/stack`
- emulating stack using the slice (RPN emulating stack)
- using a stack-like struct (RPN stack)
- using an operators list as a string (RPN operators list)
- using the same operators list and the string package (RPN operators no fields)
- emulating a turing machine (rpn (slow) Turing machine)

***How to run***

You can run examples and tests using `src>go test`.
To run benchmarks with `src>go test -bench=.`

***Results***

- Emulating the stack is the most efficient classic method.
- Using a `struct` in a provided Go Doc package is already costly by a factor of 2.
- As soon as you add slice editing, you divide performance by 3.
- The worst being to hold all elements of the expression in one slice.

Slice is the underlying type of string. So string editing won't change anything and might be
worse as a string is read-only and might be copied over and over again.

Benchmarking is done on a calculation using floats.

***Optimum***

Turing machine nears the stack emulation only if you allow Go to keep the for loop instead of re-creating it.
If you use a "while" form with increment inside the loop to follow general rules or if you exit
the for loop to re-create it, performance is dramatically degraded. 

The key is to explore []strings efficiently, i.e. rewinding the band. The key element is 
the compilation of the for loop using range. Exiting a for loop without finishing is costly.

**Results**
```
go version go1.13 windows/amd64

BenchmarkRPNEmulatingStack-4           1000000              1093 ns/op
BenchmarkRPNTuringMachine-4            1000000              1095 ns/op
BenchmarkRPNSlowTuringMachine-4        649212              1858 ns/op
BenchmarkRPNStack-4                      600396              1974 ns/op
BenchmarkRPN_operators_no_fields-4        285874              4034 ns/op
BenchmarkRPNOperatorsList-4             235432              4801 ns/op
BenchmarkRPN-4                            192178              6241 ns/op

```

 
 
 