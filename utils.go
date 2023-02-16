package sim2d

import (
	"math"
	"math/big"
)

func symmetricModulo(n *big.Rat, r int64) (*big.Int, error) {
	// Check initial radix (base) value.
	if r == 0 {
		return nil, ErrInvalidRadix
	}
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

func polynomialLength(q, p int) int {
	return q + (-1 * p) + 1
}

func expansion(polyLength, base int, numerator *big.Rat) ([]float64, error) {
	var exp []float64
	// Base as a float 64 bits.
	b := float64(base)
	for i := 0; i < polyLength; i++ {
		// Exponent.
		e := float64(i)
		// Second operand.
		nb := math.Pow(b, e)
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
func isolateNumerator(f *big.Rat, b, p int) *big.Rat {
	bp := math.Pow(float64(b), float64(-p))
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

func validateModulo(b int) error {
	// b >= 2.
	if b < 2 {
		return ErrBIsLessThan2
	}
	return nil
}

func validateSmallestPowerOfExpansion(p, q int) error {
	// p < q.
	if p >= q {
		return ErrPIsLessThanQ
	}
	// p < 0.
	if p >= 0 {
		return ErrPIsGreaterThanOrEqualToZero
	}
	return nil
}

func validateGreatestPowerOfExpansion(q int) error {
	// q > 0.
	if q <= 0 {
		return ErrQIsLessThanOrEqualToZero
	}
	return nil
}

func validateDegree(p, q, d int) error {
	// d >= 1.
	if d < 1 {
		return ErrDIsLessThanOne
	}
	// d is a power of 2.
	// Log base 2 of d.
	floatD := float64(d)
	logD := math.Log2(floatD)
	// Integer part of log base 2 of d.
	intLog := math.Round(logD)
	// Recalculated d.
	d2 := math.Pow(2.0, intLog)
	if floatD != d2 {
		return ErrDIsNotAPowerOf2
	}
	// d > q + |p|.
	absP := math.Abs(float64(p))
	if d <= (q + int(absP)) {
		return ErrDIsLessThanOrEqualToQPlusP
	}
	return nil
}

func validateParameters(b, p, q, d int) error {
	var err error
	// Validate modulo b.
	err = validateModulo(b)
	if err != nil {
		return err
	}
	// Validade smallest power of expansion.
	err = validateSmallestPowerOfExpansion(p, q)
	if err != nil {
		return err
	}
	err = validateGreatestPowerOfExpansion(q)
	if err != nil {
		return err
	}
	// Validate degree.
	err = validateDegree(p, q, d)
	if err != nil {
		return err
	}
	// Validate fraction.
	return nil
}
