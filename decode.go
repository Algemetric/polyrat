package sim2dcodec

import (
	"math/big"
)

// Decode decodes a polynomial into its original fraction.
func Decode(code []int64, b, p, q int) *big.Rat {
	// Code length.
	l := len(code)
	// decode = vector([ -1 * input[l+p+i] for i in range(-p) ] + input[:q+1])
	var original []int64
	for i := 0; i < -p; i++ {
		o := -1 * code[l+p+i]
		original = append(original, int64(o))
	}
	for i := 0; i < q+1; i++ {
		original = append(original, code[i])
	}

	xval := xVal(b, q, p)

	total := dotProduct(xval, original)

	return total
}

func xVal(b, q, p int) []*big.Rat {
	// Decoded fraction.
	var xval []*big.Rat
	// Base in 64 bits.
	b64 := int64(b)
	// Polynomial length.
	pl := polynomialLength(q, p)
	for i := 0; i < pl; i++ {
		// Fraction.
		f := big.NewRat(1, 1)

		pPlusI := int64(p + i)
		base := big.NewInt(b64)
		if pPlusI < 0 {
			e := big.NewInt(-1 * pPlusI)
			base.Exp(base, e, nil)
			f.SetFrac(base, big.NewInt(1))
			f.Inv(f)
		} else {
			e := big.NewInt(pPlusI)
			base.Exp(base, e, nil)
			f.SetInt(base)
		}
		// (1/2401, 1/343, 1/49, 1/7, 1, 7)
		xval = append(xval, f)
	}
	return xval
}
