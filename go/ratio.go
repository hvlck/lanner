// utilities for working with numbers
// mostly for arbitrary-precision math

package main

type Ratio struct {
	numerator   uint64
	denominator uint64
	positive    bool
}

// determines whether two ratios are equal to one another
// automatically simplifies the ratio passed as argument *n*
func Equals(r *Ratio, n *Ratio) bool {
	r.Simplify()
	n.Simplify()

	if (r.numerator == n.numerator) && (r.denominator == n.denominator) && (r.positive == n.positive) {
		return true
	}

	return false
}

func Multiply(r *Ratio, n *Ratio) *Ratio {
	return &Ratio{
		numerator:   r.numerator * n.numerator,
		denominator: r.denominator * n.denominator,
		positive:    r.positive == n.positive,
	}
}

func Divide(r *Ratio, n *Ratio) *Ratio {
	return &Ratio{
		numerator:   r.numerator * n.denominator,
		denominator: r.denominator * n.numerator,
		positive:    r.positive == n.positive,
	}
}

func (*Ratio) Evaluate() float64 {
	return 0
}

func (r *Ratio) Simplify() {
	factor := gcd(r.numerator, r.denominator)
	r.numerator = r.numerator / factor
	r.denominator = r.denominator / factor
}

func CreateRatio(num int64, den int64) *Ratio {
	n := float64(num)
	d := float64(den)
	p := true
	if (num < 0 && den < 0) || (num > 0 && den > 0) {
		p = false
	}

	return &Ratio{
		numerator:   uint64(Absolute(n)),
		denominator: uint64(Absolute(d)),
		positive:    p,
	}
}

// greatest common denominator
func gcd(a uint64, b uint64) uint64 {
	d := b

	c := b
	for {
		if c == 0 {
			break
		}
		d = c
		c = a % c
		a = d
	}

	return a
}
