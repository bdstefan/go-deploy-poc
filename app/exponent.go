package main

import (
	"math"
)

//Exponent struct store the base number and its exponent
type Exponent struct {
	base, exponent int
}

//Power of a number
func (ex *Exponent) power() int {
	return int(math.Pow(float64(ex.base), float64(ex.exponent)))
}
