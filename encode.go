package sim2dcodec

import (
	"math/big"
)

/*
#p: smallest power (negative)
#q: largest power (positive)
#d: degree of modulus
#base: base of Laurent expansion

Example:

[2, -3, 0, 0, -3, 0, 1, -2] = SIM2DEncode(-44979/2401,7,-4,1,8)
base = 7
p = -4
q = 1
d = 8
*/

// Sizes:
// base int
// p int
// q int
// d int

func Encode(fraction *big.Rat, base, p, q, d int) []int64 {
	// Length of the polynomial.
	polyLength := polynomialLength(q, p)
	// Numerator from the given fraction.
	numerator := fraction.Num().Int64()
	// Calculate expansion.
	exp := expansion(polyLength, base, numerator)
	// Generate encoding (code).
	code := code(p, polyLength, d, exp)
	// return encoding
	return code
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
	return c
}
