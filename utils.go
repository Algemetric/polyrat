package sim2dcodec

import (
	"math"
	"math/big"
)

func symmetricModulo(n, radix int64) (int64, error) {
	var modulo, remainder int64
	remainder = n % radix
	halfRadix := float64(-radix) / 2.0
	if remainder <= 0 && halfRadix < float64(remainder) {
		modulo = remainder
	} else {
		modulo = remainder + radix
	}
	return modulo, nil
}

func polynomialLength(q, p int) int {
	return q + (-1 * p) + 1
}

func expansion(polyLength, base int, numerator int64) ([]float64, error) {
	var exp []float64
	var baseJPlus, baseJ, c float64
	var a, b int64
	var err error
	for i := 0; i < polyLength; i++ {
		// a = SymMod(numerator, base^(j+1))
		baseJPlus = math.Pow(float64(base), float64(i+1))
		a, err = symmetricModulo(numerator, int64(baseJPlus))
		if err != nil {
			return nil, err
		}
		// b = SymMod(numerator, base^j)
		baseJ = math.Pow(float64(base), float64(i))
		b, err = symmetricModulo(numerator, int64(baseJ))
		if err != nil {
			return nil, err
		}
		c = float64(a-b) / baseJ
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
