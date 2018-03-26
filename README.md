***Reverse Polish Notation calculator***

This well known problem can be solved in Go using :
- slice editing (RPN)
- emulating stack using the slice (RPN emulating stack)
- using a stack-like struct (RPN stack)
- using an operators list as a string (RPN operators list)
- using the same operators list and the string package (RPN operators no fields)
- emulating a turing machine (rpn (slow) Turing machine)

All solutions have the same features :
- float numbers
- arithmetic operators
- binary operator ^ (exponent)
- unary operator sqrt
- panics on invalid values and operators
- input is a string

You can run examples and tests using "go test" in the src directory.
To run benchmark use "go test bench=." in the same directory.

Emulating the stack is the most efficient classic method.
Using a struct in a provided Go Doc package is already costly by a factor of 2.
As soon as you add slice editing, you divide performance by 3.
The worst being to hold all elements of the expression in one slice.

Slice is the underlying type of string. So string editing won't change anything and might be
worse as a string is read-only and might be copied over and over again.

Turing machine beats it only if you allow Go to keep the for loop instead of re-creating it.
If you use a "while" form with increment inside the loop to follow general rules or if you exit
the for loop to re-create it, performance is 20% slower than stack emulation. 

The key is to explore []strings efficiently, i.e. rewinding the band. The key element is 
the compilation of the for loop using range.

**Results**

```
 goos: windows
 goarch: amd64
 BenchmarkRPN_Turing_machine-4            2000000               761 ns/op
 BenchmarkRPN_emulating_stack-4           2000000               958 ns/op
 BenchmarkRPN_slow_Turing_machine-4       1000000              1324 ns/op
 BenchmarkRPN_stack-4                     1000000              1755 ns/op
 BenchmarkRPN_operators_no_fields-4        500000              2790 ns/op
 BenchmarkRPN_operators_list-4             500000              3176 ns/op
 BenchmarkRPN-4                            300000              3959 ns/op
   
```

 
 
 