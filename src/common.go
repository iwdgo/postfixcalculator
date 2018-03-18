package RPN

//This avoids various conflicts in the package
var err error

//Two examples and there expected results
const RPNInput = "6 8 4.0 sqrt + 3.01 1.99 + - *" //= 30.000000
const input = "3 4 2 * 1 5 - 2 3 ^ ^ / +"         //= 3.0001220703125
