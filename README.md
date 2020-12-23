[![Go Reference](https://pkg.go.dev/badge/iwdgo/postfixcalculator.svg)](https://pkg.go.dev/iwdgo/postfixcalculator)
[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/postfixcalculator)](https://goreportcard.com/report/github.com/iwdgo/postfixcalculator)
[![codecov](https://codecov.io/gh/iWdGo/postfixcalculator/branch/master/graph/badge.svg)](https://codecov.io/gh/iWdGo/postfixcalculator)

# Reverse Polish Notation calculator

Several solutions to this well-known problem show dramatic differences depending on the approach.
A Turing machine emulation beats all of them.

## Features

All solutions support :
- float numbers
- arithmetic operators
- binary operator ^ (exponent)
- unary operator sqrt

Input is a string. Process panics on invalid values and operators.

## Solutions

- [slice editing](https://pkg.go.dev/github.com/iwdgo/stack)
- [emulating stack using the slice](https://pkg.go.dev/github.com/iwdgo/emulatingstack)
- [using a stack-like struct](https://pkg.go.dev/github.com/iwdgo/emulatingstack)
- [using string to hold a list of operators in two loops](https://pkg.go.dev/github.com/iwdgo/operatorslist)
- [using string to hold a list of operators in one loop](https://pkg.go.dev/github.com/operatorsnofields)
- Turing machine:
  - [slow](https://pkg.go.dev/github.com/iwdgo/slowturingmachine)
  - [fast](https://pkg.go.dev/github.com/iwdgo/turingmachine)

## How to

To run tests: `>go test ./...`
To run benchmarks `>go test -bench=. ./...`

## Results

- Emulating the stack is the most efficient classic method.
- Using a `struct` from an existing package is already costly by a factor of 2.
- Adding slice editing slows performance by a factor of 3.
- The worst occurs when all elements of the expression are in one slice.

Slice is the underlying type of string. So string editing won't change anything and might be
worse as a string is read-only and is copied many times.

## Optimum

Turing machine nears the stack emulation only if you allow Go to keep the for loop instead of re-creating it.
Performance degrades dramatically when using a "while" form with increment inside the loop to follow usual
pracice or if the `for` loop is re-created. 

The key is to explore []strings efficiently, i.e. rewinding the band. The key element is 
the compilation of the for loop using range. Exiting a `for` loop without finishing is also costly.

## Benchmark

Benchmarking is done on a calculation using floats.

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

 
 
 