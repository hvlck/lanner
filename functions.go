package main

import (
	"errors"
	"math"
)

// todo: support floating point
// very simple and unoptimized implementation
// currently only supports whole numbers
func Factorial(n uint64) uint64 {
	p := 1
	for idx := 1; idx <= int(n); idx++ {
		p = p * idx
	}

	return uint64(p)
}
// raise a number **base** to a power **pow**
func Pow(base float64, pow float64) float64 {
	return math.Pow(base, pow)
}

// vectors and vector resolution
func Vec() {

}

func Absolute(num float64) float64 {
	if math.IsNaN(num) {
		return math.NaN()
	}

	if num < 0 {
		return num * -1
	}

	return num
}
