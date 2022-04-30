// utilities for working with numbers
// mostly for arbitrary-precision math

package ratio

import "errors"

type Ratio struct {
	numerator   int64
	denominator int64
}

// determines whether two ratios are equal to one another
// automatically simplifies the ratio passed as argument *n*
func Equals(r *Ratio, n *Ratio) bool {
	r.Simplify()
	n.Simplify()

	if (r.numerator == n.numerator) && (r.denominator == n.denominator) {
		return true
	}

	return false
}

func Add(r *Ratio, n *Ratio) *Ratio {
	f := Ratio{
		numerator:   (r.numerator * n.denominator) + (n.numerator * r.denominator),
		denominator: r.denominator * n.denominator,
	}

	f.Simplify()

	return &f
}

func Subtract(r *Ratio, n *Ratio) *Ratio {
	f := Ratio{
		numerator:   (r.numerator * n.denominator) - (n.numerator * r.denominator),
		denominator: r.denominator * n.denominator,
	}

	f.Simplify()

	return &f
}

func Multiply(r *Ratio, n *Ratio) *Ratio {
	f := Ratio{
		numerator:   r.numerator * n.numerator,
		denominator: r.denominator * n.denominator,
	}

	f.Simplify()

	return &f
}

func Divide(r *Ratio, n *Ratio) *Ratio {
	f := Ratio{
		numerator:   r.numerator * n.denominator,
		denominator: r.denominator * n.numerator,
	}

	f.Simplify()

	return &f
}

func (r *Ratio) Evaluate() (float64, error) {
	if r.denominator == 0 {
		return 0, errors.New("division by zero error")
	}

	i := r.numerator / r.denominator
	return float64(((r.numerator - i) * r.denominator) / r.denominator), nil
}

func (r *Ratio) Simplify() {
	factor := int64(gcd(r.numerator, r.denominator))
	r.numerator = r.numerator / factor
	r.denominator = r.denominator / factor
}

func CreateRatio(num int64, den int64) *Ratio {
	n := float64(num)
	d := float64(den)

	return &Ratio{
		numerator:   int64(n),
		denominator: int64(d),
	}
}

// greatest common denominator
func gcd(a int64, b int64) uint64 {
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

	return uint64(a)
}
