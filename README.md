# lanner

math evaluator

Primary target is WebAssembly, but should work as a Rust crate as well.

## Syntax

Anything wrapped in parantheses (`(expr)`) or pipes (`|expr|`) that evaluates to a single value is valid.

### Constants

+ `E` - [Euler's number](https://en.wikipedia.org/wiki/E_(mathematical_constant))
+ `i` - [Imaginary unit](https://en.wikipedia.org/wiki/Imaginary_unit)
+ `Tau` - [Tau (2*PI)](https://en.wikipedia.org/wiki/Turn_(angle)#Tau_proposals)
+ `PI` - [Pi](https://en.wikipedia.org/wiki/Pi)

```lanner
E + 10
Tau * 40
Pi - 3.14
```

### Functions

+ `sin(x)` - sine
+ `cos(x)` - cosine
+ `tan(x)` - tangent
+ `asin(x)` - arcsin
+ `acos(x)` - arccosine
+ `atan(x)` - arctangent
+ `sqrt(x)` - square root
+ `abs(x)` - absolute value

### Operators

+ `|expr|` - returns the absolute value of the given expression
+ `int + int` - returns the sum of two numbers
+ `int - int` - returns the difference of two numbers
+ `int * int` - returns the product of two numbers
+ `int / int` - returns the quotient of two numbers
+ `int - int` - returns the difference of two numbers
+ `int % int` - performs a [modulo operation](https://en.wikipedia.org/wiki/Modulo_operation) on two numbers
+ `num ^ power` - raises a number to the given power
+ (not implemented) `exponent ^/ int` - returns the nth (exponent) roots of the given int

## Todo

+ make functions/conversions able to fit in the Order of Operations hierarchy evaluation

## Roadmap

+ equality checks (gt/lt/equal)
+ unit conversion
+ variables/substitution
+ trigonometry
+ calculus
+ linear algebra
+ constants
+ engineering
+ compile-time math macro?

+ functions
  + nthRoot
  + nth

## Known Issues

+ order of operations doesn't work
+ nesting values doesn't work
