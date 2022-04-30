package main

import (
	"math"
)

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
