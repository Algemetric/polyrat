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

	// xval = vector([base^(p+i) for i in range(q-p+1)])
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
		// TODO: find a trustworthy way.
		pPlusI := int64(p + i)
		if pPlusI < 0 {
			e := big.NewInt(int64(-1 * pPlusI))
			base := big.NewInt(b64)
			base.Exp(base, e, nil)
			f.SetFrac(base, big.NewInt(1))
			f.Inv(f)
		} else {
			base := big.NewInt(b64)
			base.Exp(base, big.NewInt(pPlusI), nil)
			f.SetInt(base)
		}
		// (1/2401, 1/343, 1/49, 1/7, 1, 7)
		xval = append(xval, f)
	}
	return xval
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
