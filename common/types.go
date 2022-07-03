package common

import "fmt"

// any value that can be solved for
type Determinate struct {
	Src   interface{}
	Value uint
}

// returns the value of the determinate
// todo: functional implementation
func (d *Determinate) Resolve() string {
	return fmt.Sprint(d.Value)
}
