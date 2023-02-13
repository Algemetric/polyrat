package sim2dcodec

import (
	"math"
	"math/big"
)

func symmetricModulo(n, radix int64) (int64, error) {
	var modulo, remainder int64
	// Check radix value.
	if radix == 0 {
		return modulo, ErrInvalidRadix
	}
	remainder = n % radix

	modulo = remainder + radix
	halfRadix := float64(-radix) / 2.0
	if remainder <= 0 && halfRadix < float64(remainder) {
		modulo -= radix
	}
	return modulo, nil
}

func polynomialLength(q, p int) int {
	return q + (-1 * p) + 1
}

func expansion(polyLength, base int, numerator int64) ([]float64, error) {
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
		nb *= b
		fo, err := symmetricModulo(numerator, int64(nb))
		if err != nil {
			return nil, err
		}
		c := float64(fo-so) / (nb / b)
		exp = append(exp, c)
	}
	return exp, nil
}

func dotProduct(v1 []*big.Rat, v2 []int64) *big.Rat {
	// Dot product total.
	t := big.NewRat(0, 1)
	for i := 0; i < len(v2); i++ {
		// Multiplication step.
		f := big.NewRat(v2[i], 1)
		f.Mul(f, v1[i])
		// Addition step.
		t.Add(t, f)
	}
	return t
}
