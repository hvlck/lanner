# lanner

arbitrary-precision math evaluator

Primary target is WebAssembly, but should work as a go package as well.

--------------WARNING: VERY MUCH A WORK IN PROGRESS---------------
DO NOT EXPECT VALUES TO BE ARBITRARILY ACCURATE OR FOR THE PACKAGE TO WORK AS INTENDED JUST YET.

## Priorities

+ accuracy
+ speed
+ actually implementing things

## Syntax

Anything wrapped in parentheses (`(expr)`) or pipes (`|expr|`) that evaluates to a single value is valid.

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
+ `csx(x)` - cosecant
+ `sec(x)` - secant
+ `cot(x)` - cotangent
+ `log_b(x)` - base-*b* logarithm

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
  + `mi` or `mile`
+ variables/substitution
+ trigonometry
+ calculus
+ linear algebra
+ constants
+ engineering
  + physics
    + kinematics
  + civil - structural properties, etc.
  + thermodynamics and chemistry
+ compile-time math macro?
+ arbitrary-precision math
+ comment and documentation system
+ finish language syntax
+ multiline parsing
+ logarithms
+ factorials with `x!`
+ formal specification
+ statistics
+ graphing backend
+ matrix and vector ops
+ financials
+ boolean algebra
  + `==`
  + `!=`
  + etc.

+ functions
  + nthRoot
  + nth

## Known Issues

+ order of operations doesn't work
+ nesting values doesn't work

## Constants Reference

+ [The On-Line Encyclopedia of Integer Sequences® (OEIS®)](https://oeis.org/)

## Inspiration and Thanks

+ [GitHub - p-e-w/savage: A primitive computer algebra system](https://github.com/p-e-w/savage)
