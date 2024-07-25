package pkg

import "math"

// NewOptimus opt := NewOptimus(2123809381, 1885413229, 146808189, 31)
func NewOptimus(prime, inverse, random int, bitSize int) *Optimus {
	// https://github.com/jenssegers/optimus
	m := int(math.Pow(2, float64(bitSize)) - 1)
	return &Optimus{
		prime:   prime,
		inverse: inverse,
		xor:     random,
		max:     m,
	}
}

type Optimus struct {
	prime   int
	inverse int
	xor     int
	max     int
}

func (op *Optimus) Encode(n int) int {
	return ((n * op.prime) & op.max) ^ op.xor
}

func (op *Optimus) Decode(n int) int {

	return ((n ^ op.xor) * op.inverse) & op.max
}
