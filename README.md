[![Go Reference](https://pkg.go.dev/badge/github.com/iwdgo/postfixcalculator.svg)](https://pkg.go.dev/iwdgo/postfixcalculator)
[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/postfixcalculator)](https://goreportcard.com/report/github.com/iwdgo/postfixcalculator)
[![codecov](https://codecov.io/gh/iwdgo/postfixcalculator/branch/master/graph/badge.svg)](https://codecov.io/gh/iWdGo/postfixcalculator)

[![Build status](https://ci.appveyor.com/api/projects/status/pnnlu9oovyo71d6q?svg=true)](https://ci.appveyor.com/project/iWdGo/postfixcalculator)
![GitHub](https://github.com/iwdgo/postfixcalculator/workflows/GitHub/badge.svg)

# Reverse Polish Notation calculator

Several solutions to this well-known problem show dramatic differences depending on the approach.
A Turing machine emulation rivals the method using the stack emulation of reference.

## Features

All solutions support :
- float numbers
- arithmetic operators
- binary operator ^ (exponent)
- unary operator sqrt

Input is a string where elements are spaced. Process panics on invalid values and operators.
Float values are treated as 64 bits.
The original expression is used to store values and edited. When it contains only one element,
calculation is complete and the result is returned as a `float64`.

## Solutions

- [slice editing](https://pkg.go.dev/github.com/iwdgo/postfixcalculator/stack)
- [emulating stack using the slice](https://pkg.go.dev/github.com/iwdgo/postfixcalculator/emulatingstack)
- [using a stack-like struct](https://pkg.go.dev/github.com/iwdgo/postfixcalculator/emulatingstack)
- [using string to hold a list of operators in two loops](https://pkg.go.dev/github.com/iwdgo/postfixcalculator/operatorslist)
- [using string to hold a list of operators in one loop](https://pkg.go.dev/github.com/iwdgo/postfixcalculator/operatorsnofields)
- Turing machine:
  - [slow](https://pkg.go.dev/github.com/iwdgo/postfixcalculator/slowturingmachine)
  - [fast](https://pkg.go.dev/github.com/iwdgo/postfixcalculator/turingmachine)

## How to

There are examples in every package. A first use that prints 4 could be
```
package main

import "github.com/iwdgo/postfixcalculator/turingmachine"

func main() {
	println(turingmachine.RPNTuringMachine("2 2 +"))
}

```


## Results

- Emulating the stack is the most efficient classic method.
- Using a `struct` from an existing package divides performance by 2.
- Adding slice editing slows performance by a factor of 3.
- The worst occurs when all elements of the expression are in one slice.

Slice is the underlying type of string. So string editing will probably not improve if not worsen 
as a string is read-only and could be copied many times during the process.

## Optimum

Turing machine nears the stack emulation only if Go re-uses the `for` loop instead of re-creating it.
Performance degrades dramatically when using a "while" form with increment inside the loop to follow usual
practice.

The key is to explore `[]strings` efficiently, i.e. rewinding the band. The key element is 
the compilation of the `for` loop using range. 
Exiting a `for` loop without finishing is also costly.

## Tests

To run tests: `>go test ./...`

## Benchmark

To run benchmarks `>go test -bench=. ./...`
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
