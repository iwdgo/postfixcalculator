***Reverse Polish Notation calculator***

This well known problem can be solved in Go using :
- slice editing (RPN)
- emulating stack using the slice (RPN emulating stack)
- using a stack-like struct (RPN stack)
- using an operators list as a string (RPN operators list)
- using the same operators list and the string package (RPN operators no fields)
- emulating a turing machine (rpn Turing machine)

All solutions have the same features :
- float numbers
- arithmetic operators
- binary operator ^ (exponent)
- unary operator sqrt
- panics on invalid values and operators
- input is a string

You can run examples and tests using "go test" in the src directory.
To run benchmark use "go test bench=." in the same directory.

Emulating the stack is the most efficient method.
Using a struct in a provided Go Doc package is already costly by a factor of 2.
As soon as you add slice editing, you divide performance by 3.
The worst being to hold all elements of the expression in one slice.
Slice is the underlying type of string. So string editing won't change anything and might be
worse as a string is read-only and might be copied over and over again.

Turing machine beats it only if you allow Go to keep the for loop instead of re-creating it
Otherwise, it loses 20% on the exploration of the []strings which is rewinding the band.

**Results**

```
 goos: windows
 goarch: amd64
 BenchmarkRPN_Turing_machine-4            2000000               774 ns/op
 BenchmarkRPN_emulating_stack-4           2000000               907 ns/op
 BenchmarkRPN_stack-4                     1000000              1693 ns/op
 BenchmarkRPN_operators_no_fields-4        500000              2750 ns/op
 BenchmarkRPN_operators_list-4             500000              3111 ns/op
 BenchmarkRPN-4                            500000              3780 ns/op
 
```

 
 
 