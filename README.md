# Reverse Polish Notation calculator

Several solution to this well-known problem show dramatic differences depending on the approach.
A Turing machine emulation beats all of them.

***Features***

All solutions support :
- float numbers
- arithmetic operators
- binary operator ^ (exponent)
- unary operator sqrt

Input is a string. Panic on invalid values and operators

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
To run benchmarks with `src>go test bench=."`

***Results***

- Emulating the stack is the most efficient classic method.
- Using a `struct` in a provided Go Doc package is already costly by a factor of 2.
- As soon as you add slice editing, you divide performance by 3.
- The worst being to hold all elements of the expression in one slice.

Slice is the underlying type of string. So string editing won't change anything and might be
worse as a string is read-only and might be copied over and over again.

Benchmarking is done on a calculation where floats appear.

***Optimum***

Turing machine nears the stack emulation only if you allow Go to keep the for loop instead of re-creating it.
If you use a "while" form with increment inside the loop to follow general rules or if you exit
the for loop to re-create it, performance is dramatically degraded. 

The key is to explore []strings efficiently, i.e. rewinding the band. The key element is 
the compilation of the for loop using range. Exiting a for loop without finishing is costly.

**Results**
```
go version go1.11.2 windows/amd64

BenchmarkRPN_emulating_stack-4           2000000               945 ns/op
BenchmarkRPN_Turing_machine-4            1000000              1020 ns/op
BenchmarkRPN_stack-4                     1000000              1811 ns/op
BenchmarkRPN_slow_Turing_machine-4       1000000              2203 ns/op
BenchmarkRPN_operators_no_fields-4        500000              3643 ns/op
BenchmarkRPN_operators_list-4             300000              4260 ns/op
BenchmarkRPN-4                            300000              5248 ns/op

```

 
 
 