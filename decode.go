package sim2d

import (
	"math/big"
)

// Decode decodes a polynomial into its original fraction.
func Decode(code []int64, b, p, q int) (*big.Rat, error) {
	// Validate input.
	err := validateDecodingParameters(code, b, p, q)
	if err != nil {
		return nil, err
	}
	// Code length.
	l := len(code)
	var original []int64
	for i := 0; i < -p; i++ {
		o := -1 * code[l+p+i]
		original = append(original, int64(o))
	}
	for i := 0; i < q+1; i++ {
		original = append(original, code[i])
	}
	// Decoding powers used for evaluation.
	ep := evaluationPowers(b, q, p)
	// Return fraction.
	return dotProduct(ep, original), nil
}

func evaluationPowers(b, q, p int) []*big.Rat {
	// Fractions.
	var powers []*big.Rat
	// Polynomial length.
	pl := polynomialLength(q, p)
	for i := 0; i < pl; i++ {
		// Fraction.
		f := big.NewRat(1, 1)
		pPlusI := int64(p + i)
		base := big.NewInt(int64(b))
		if pPlusI < 0 {
			// Define fractions.
			e := big.NewInt(-1 * pPlusI)
			base.Exp(base, e, nil)
			f.SetInt(base)
			// Invert fraction.
			f.Inv(f)
		} else {
			// Define fractions.
			e := big.NewInt(pPlusI)
			base.Exp(base, e, nil)
			f.SetInt(base)
		}
		// Get evaluation powers.
		powers = append(powers, f)
	}
	return powers
}
