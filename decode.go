package sim2d

import (
	"math/big"
)

// Decode decodes a polynomial into its original fraction.
func Decode(code []int64, params *Parameters) (*big.Rat, error) {
	// Validate input.
	err := validateDecodingParameters(code, params)
	if err != nil {
		return nil, err
	}
	// Code length.
	l := len(code)
	var original []int64
	for i := 0; i < -params.P; i++ {
		o := -1 * code[l+params.P+i]
		original = append(original, int64(o))
	}
	for i := 0; i < params.Q+1; i++ {
		original = append(original, code[i])
	}
	// Decoding powers used for evaluation.
	ep := evaluationPowers(params)
	// Return fraction.
	return dotProduct(ep, original), nil
}

func evaluationPowers(params *Parameters) []*big.Rat {
	// Fractions.
	var powers []*big.Rat
	// Polynomial length.
	pl := polynomialLength(params)
	for i := 0; i < pl; i++ {
		// Fraction.
		f := big.NewRat(1, 1)
		pPlusI := int64(params.P + i)
		base := big.NewInt(int64(params.B))
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
