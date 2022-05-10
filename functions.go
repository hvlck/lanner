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

// todo: memoize for special angles

// func Asin(n int8) (float64, error) {
// 	if n > math.Pi/2 || n < -(math.Pi/2) {
// 		return 0, errors.New("out of domain")
// 	}

// 	// use pythagorean identity

// 	return 0.0, nil
// }

func Acos(n int8) (float64, error) {
	if n > 1 || n < 1 {
		return 0, errors.New("out of domain")
	}

	return 0, nil
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
