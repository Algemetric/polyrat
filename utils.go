package sim2d

import (
	"math"
	"math/big"
)

func symmetricModulo(n *big.Rat, r int64) (*big.Int, error) {
	// Radix big integer.
	radix := big.NewInt(r)
	// Modulo = numerator % radix.
	// We know that the numerator went through a process to produce n/1,
	// therefore we can just take the numerator to calculate the modulo.
	modulo := big.NewInt(0)
	modulo.Mod(n.Num(), radix)
	// Half radix.
	// halfRadix = radix / 2
	halfRadix := big.NewRat(1, 2)
	rationalRadix := big.NewRat(1, 1)
	rationalRadix.SetInt(radix)
	halfRadix.Mul(halfRadix, rationalRadix)
	// remainder <= 0 && halfRadix < remainder.
	// To execute "0 <= modulo" we need to do "0 < modulo" and "0 == modulo" separately.
	// Therefore, to have better visualization of the logic, we will give names for the comparison operands.
	zero := big.NewRat(0, 1)
	rationalModulo := big.NewRat(1, 1)
	rationalModulo.SetInt(modulo)
	moduloIsGreaterThanOrEqualToZero := zero.Cmp(rationalModulo) == -1 || zero.Cmp(rationalModulo) == 0
	moduloIsLessThanHalfRadix := halfRadix.Cmp(rationalModulo) == -1
	// remainder <= 0 && halfRadix < remainder.
	if moduloIsGreaterThanOrEqualToZero && moduloIsLessThanHalfRadix {
		rationalModulo.Sub(rationalModulo, rationalRadix)
	}

	return rationalModulo.Num(), nil
}

func polynomialLength(params *Parameters) int {
	return params.Q + (-1 * params.P) + 1
}

func expansion(numerator *big.Rat, params *Parameters) ([]float64, error) {
	var exp []float64
	// Length of the polynomial.
	pl := polynomialLength(params)
	// Base as a float 64 bits.
	b := float64(params.B)
	for i := 0; i < pl; i++ {
		// Second operand.
		nb := math.Pow(b, float64(i))
		so, err := symmetricModulo(numerator, int64(nb))
		if err != nil {
			return nil, err
		}
		// First operand.
		// b^(e+1).
		fo, err := symmetricModulo(numerator, int64(nb*b))
		if err != nil {
			return nil, err
		}
		// (fo-so) / (nb / b)
		fo.Sub(fo, so)
		c := big.NewRat(fo.Int64(), int64(nb))
		c2, _ := c.Float64()

		exp = append(exp, c2)
	}
	return exp, nil
}

// isolateNumerator makes the fraction display the numerator over 1.
func isolateNumerator(f *big.Rat, params *Parameters) *big.Rat {
	bp := math.Pow(float64(params.B), float64(-params.P))
	db := big.NewRat(int64(bp), 1)
	return f.Mul(f, db)
}

func dotProduct(v1 []*big.Rat, v2 []int64) *big.Rat {
	// Dot product total.
	dp := big.NewRat(0, 1)
	// Fraction to represent the multiplication of terms.
	f := big.NewRat(1, 1)
	for i := 0; i < len(v2); i++ {
		// Multiplication step.
		f.SetInt64(v2[i])
		f.Mul(f, v1[i])
		// Addition step.
		dp.Add(dp, f)
	}
	return dp
}
