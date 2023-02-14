package sim2dcodec

import (
	"fmt"
	"math"
	"math/big"
)

func symmetricModulo(n *big.Rat, r int64) (*big.Int, error) {
	// Check initial radix value.
	if r == 0 {
		return nil, ErrInvalidRadix
	}
	radix := big.NewInt(r)
	// Modulo: numerator % radix.
	remainder := big.NewInt(0)
	remainder.Mod(n.Num(), radix)
	fmt.Printf("\n%s mod %s = %s\n", n.Num().String(), radix.String(), remainder.String())
	// Big int addition.
	// modulo := remainder + radix
	modulo := big.NewInt(0)
	modulo.Add(modulo, remainder)
	modulo.Add(modulo, radix)
	// Half radix.
	// halfRadix := -radix / 2.0
	halfRadix := big.NewRat(-radix.Int64(), 2)
	fmt.Printf("\nmodulo = %s, halfRadix = %s\n", modulo.String(), halfRadix.String())

	// remainder <= 0 && halfRadix < remainder.
	// To execute "<=0" we need to do "<0" and "==0" separately.
	zero := big.NewInt(0)
	remainderFraction := big.NewRat(remainder.Int64(), 1)
	if (remainder.Cmp(zero) == -1 || remainder.Cmp(zero) == 0) && (halfRadix.Cmp(remainderFraction) == -1) {
		modulo.Sub(modulo, radix)
	}

	return modulo, nil
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
		fmt.Printf("\nc = %s", c.String())
		exp = append(exp, c2)
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
