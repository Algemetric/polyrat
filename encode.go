package sim2d

import (
	"math/big"
)

// Encode encodes a fraction into a polynomial.
func Encode(fraction *big.Rat, b, p, q, d int) ([]int64, error) {
	// Length of the polynomial.
	pl := polynomialLength(q, p)
	// Numerator from the given fraction.
	n := isolateNumerator(fraction, b, p)
	// Calculate expansion.
	e, err := expansion(pl, b, n)
	if err != nil {
		return nil, err
	}
	// Generate encoding (code).
	c := code(p, pl, d, e)
	// return encoding
	return c, nil
}

func code(p, l, d int, exp []float64) []int64 {
	var c []int64
	// Members of the expansion from -p to the end of the vector.
	for i := -p; i < len(exp); i++ {
		c = append(c, int64(exp[i]))
	}
	// d - l gives the number of zeros to be concatenated at the vector.
	zeros := d - l
	for i := 0; i < zeros; i++ {
		c = append(c, 0)
	}
	// Concatenate to the vector, from 0 to (-p) - 1, the elements of the expansion.
	for i := 0; i < -p; i++ {
		c = append(c, -int64(exp[i]))
	}
	// Return code.
	return c
}
