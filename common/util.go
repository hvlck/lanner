package common

import (
	"errors"
	ratio "lanner/common/ratio"
	"math"
)

func ToPrecision(n float64, p uint64) float64 {
	return math.Round(n*math.Pow(10, float64(p))) / math.Pow(10, float64(p))
}

// long-divides a number to the specified precision
func Divide(dividend float64, divisor float64, precision uint64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("division by zero error")
	}

	f := ToPrecision(dividend/divisor, precision)

	return f, nil
}

// converts the given value into a Ratio
func ToRatio(o interface{}) interface{} {
	switch t := o.(type) {
	case ratio.Ratio:
		return o
	case int64:
		return t
	}

	return nil
}

// converts the given value into an integer
// func toInt(o)

// determines if the value of *l* is less than *r*
func Lt(l interface{}, r interface{}) bool {
	l = ToRatio(l).(ratio.Ratio)
	r = ToRatio(r).(ratio.Ratio)

	return false
}

// sums values
func Sum(d ...Determinate) float64 {
	r := 0.0
	for _, i := range d {
		v := i.resolve()
		r += v
	}

	return r
}
