package common

import (
	"errors"
	"math"
)

func toPrecision(n float64, p uint64) float64 {
	return math.Round(n*math.Pow(10, float64(p))) / math.Pow(10, float64(p))
}

// long-divides a number to the specified precision
func Divide(dividend float64, divisor float64, precision uint64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("division by zero error")
	}

	f := toPrecision(dividend/divisor, precision)

	return f, nil
}

// converts the given value into a Ratio
func toRatio(o interface{}) interface{} {
	switch t := o.(type) {
	case Ratio:
		return o
	case int64:
		return t
	}

	return nil
}

// converts the given value into an integer
// func toInt(o)

// determines if the value of *l* is less than *r*
func lt(l interface{}, r interface{}) bool {
	l = toRatio(l).(Ratio)
	r = toRatio(r).(Ratio)

	return false
}
