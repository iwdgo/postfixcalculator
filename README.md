[![Go Reference](https://pkg.go.dev/badge/github.com/iwdgo/postfixcalculator.svg)](https://pkg.go.dev/github.com/iwdgo/postfixcalculator)
[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/postfixcalculator)](https://goreportcard.com/report/github.com/iwdgo/postfixcalculator)
[![codecov](https://codecov.io/gh/iwdgo/postfixcalculator/branch/master/graph/badge.svg)](https://codecov.io/gh/iWdGo/postfixcalculator)

[![Build status](https://ci.appveyor.com/api/projects/status/pnnlu9oovyo71d6q?svg=true)](https://ci.appveyor.com/project/iWdGo/postfixcalculator)
[![Go](https://github.com/iwdgo/postfixcalculator/actions/workflows/go.yml/badge.svg)](https://github.com/iwdgo/postfixcalculator/actions/workflows/go.yml)

# Reverse Polish Notation calculator

Several solutions to this well-known problem show dramatic differences depending on the approach.
A Turing machine emulation rivals the method using the stack emulation of reference.

## Features

All solutions support :
- float numbers
- arithmetic operators written + - * /
- binary operator ^ (exponent)
- unary operator `sqrt`

Input is a string where elements are spaced. Calculation panics on invalid values and operators.
Float values are parsed as 64 bits.
The original expression is used to store values and edited. When one result is left,
calculation is complete and the result is returned.

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

There are examples in every package. A first use that prints 4 could be:

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
- The worst occurs when all elements of the expression and its results are in one slice as
conversion to and from numbers is an expensive operation.

Slice is the underlying type of string. So string editing will probably not improve if not worsen 
as a string is read-only and could be copied many times during the process.

## Optimum

Turing machine beats the stack emulation only if the `for` loop is re-used instead of re-creating it.
Performance degrades dramatically when using a `while` form with increment inside the loop to follow usual
practice.

Exploring the slice of strings containing all terms of the expression efficiently is executed by rewinding the band. 
Initializing and exiting without finishing a `for` loop using `range` is costly.

## Evaluation

To run tests: `>go test ./...`

To run benchmarks `>go test -bench=. ./...`

```
go version go1.19.4 windows/amd64

$ go test -bench=. -benchtime=5s ./... | grep Benchmark
BenchmarkRPNEmulatingStack-4          5973022              1006 ns/op
BenchmarkRPN-4                         963788              5985 ns/op
BenchmarkRPNOperatorsList-4           1248453              4903 ns/op
BenchmarkRPNOperatorsNoFields-4       1488712              3969 ns/op
BenchmarkRPNSlowTuringMachine-4       3583735              1690 ns/op
BenchmarkRPNStack-4                   3169281              1886 ns/op
BenchmarkRPNTuringMachine-4           6172328               981.0 ns/op

```
