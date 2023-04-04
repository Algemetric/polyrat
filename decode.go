package sim2d

import (
	"math/big"
)

// Decode decodes a polynomial into its original rational.
func Decode(code []int64, params *Parameters) (float64, error) {
	// Validate input.
	err := validateDecodingParameters(code, params)
	if err != nil {
		return 0.0, err
	}
	// Code length.
	l := len(code)
	var original []int64
	for i := 0; i < -params.MinPower(); i++ {
		o := -1 * code[l+params.MinPower()+i]
		original = append(original, int64(o))
	}
	for i := 0; i < params.MaxPower()+1; i++ {
		original = append(original, code[i])
	}
	// Decoding powers used for evaluation.
	ep := evaluationPowers(params)
	// Fraction.
	f := dotProduct(ep, original)
	// Calculates rational from fraction with "exact" flag.
	r, e := f.Float64()
	// If rational was not exact, then round it.
	// TODO: check rationals ending in > 5 and < 5.
	if !e {
		r = roundUp(r, params)
	}
	return r, nil
}

func evaluationPowers(params *Parameters) []*big.Rat {
	// Fractions.
	var powers []*big.Rat
	// Polynomial length.
	pl := polynomialLength(params)
	for i := 0; i < pl; i++ {
		// Fraction.
		f := big.NewRat(1, 1)
		pPlusI := int64(params.MinPower() + i)
		base := big.NewInt(int64(params.Base()))
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
