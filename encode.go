package sim2d

import "math/big"

// Encode encodes a rational number into a set of polynomial degrees.
// The function accepts as inputs a fraction and a set of parameters.
// This fraction can be given by a 64-bits numerator, and denominator,
// or by a string representing the fraction at hand.
// The string input is more suitable for arbitrary-size numbers and will
// be transformed into an irreducible version of the same fraction before
// calculation.
// Keeping rationals as fractions during computations will preserve accuracy
// in the generation of expansions and the final set of polynomial powers.
func Encode(num, den *big.Int, b, p, q, d int) ([]int64, error) {
	// Input validation.
	// All the parameters will be analyzed through specific bounds
	// and also by the relationships between them.
	err := validateEncodingParameters(num, den, b, p, q, d)
	if err != nil {
		return nil, err
	}
	// Generate fraction with the given numerator and denominator.
	f := big.NewRat(1, 1)
	f.SetFrac(num, den)
	// Length of the polynomial.
	pl := polynomialLength(q, p)
	// Numerator from the given fraction.
	n := isolateNumerator(f, b, p)
	// Calculate expansion.
	e, err := expansion(pl, b, n)
	if err != nil {
		return nil, err
	}
	// Generate encoding (code).
	c := generateCode(p, pl, d, e)
	// return encoding.
	return c, nil
}

func generateCode(p, l, d int, exp []float64) []int64 {
	var code []int64
	// Members of the expansion from -p to the end of the vector.
	for i := -p; i < len(exp); i++ {
		code = append(code, int64(exp[i]))
	}
	// d - l gives the number of zeros to be concatenated at the vector.
	for i := 0; i < (d - l); i++ {
		code = append(code, 0)
	}
	// Concatenate to the vector, from 0 to (-p) - 1, the elements of the expansion.
	for i := 0; i < -p; i++ {
		code = append(code, -int64(exp[i]))
	}
	// Return code.
	return code
}
