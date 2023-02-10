package sim2dcodec

import (
	"math/big"
)

/*

#p: smallest power (negative)
#q: largest power (positive)
#d: degree of modulus
#base: base of Laurent expansion

def SIM2DDecode(input,base,p,q):
  l = len(input)
  decode = vector([-1*input[l+p+i]for i in range(-p)]+input[:q+1])
  xval = vector([base^(p+i) for i in range(q-p+1)])
  fraction = decode*xval
  return fraction

Example:

-44979/2401 = SIM2DDecode([2, -3, 0, 0, -3, 0, 1, -2],7,-4,1)

*/

func Decode(code []int64, base, p, q int) *big.Rat {
	// l = len(input)
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
	var xval []*big.Rat
	// var xval []float64
	count := q + (-1 * p) + 1
	for i := 0; i < count; i++ {
		frac := big.NewRat(1, 1)
		// TODO: find a trustworthy way.
		if p+i < 0 {
			e := big.NewInt(int64(-1 * (p + i)))
			basePow := big.NewInt(int64(base))
			basePow.Exp(basePow, e, nil)
			frac.SetFrac(basePow, big.NewInt(1))
			frac.Inv(frac)
		} else {
			basePow := big.NewInt(int64(base))
			basePow.Exp(basePow, big.NewInt(int64(p+i)), nil)
			frac.SetInt(basePow)
		}
		// (1/2401, 1/343, 1/49, 1/7, 1, 7)
		xval = append(xval, frac)
	}

	// fraction = decode*xval
	// This means a dot product in Sage.
	total := big.NewRat(0, 1)
	for i := 0; i < len(original); i++ {
		// Multiplication step.
		f := big.NewRat(original[i], 1)
		f.Mul(f, xval[i])
		// Addition step.
		total.Add(total, f)
	}

	// -44979/2401 = -18.7334443982
	return total
}
