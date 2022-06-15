package common

// any value that can be solved for
type Determinate struct {
	src   interface{}
	value float64
}

// returns the value of the determinate
// todo: functional implementation
func (d *Determinate) resolve() float64 {
	return 0
}
