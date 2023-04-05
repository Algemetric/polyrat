package polyrat

import (
	"math"
	"math/big"
)

func symmetricModulo(n int64) int64 {
	// Base.
	b := int64(Base)
	// Remainder.
	r := n % b
	// This condition is the inverse of 0 <= r && r < m/2.
	if r < 0 || b/2 <= r {
		r -= b
	}
	return r
}

func polynomialLength(params *Parameters) int {
	return params.MaxPower() + (-1 * params.MinPower()) + 1
}

func expansion(numerator int64, params *Parameters) []int64 {
	var exp []int64
	// Length of the polynomial.
	pl := polynomialLength(params)
	// Base.
	b := int64(Base)
	// First denominator of the sequence (d^0=1, d^1=10, d^2=100, ...).
	d := int64(1)
	for i := 0; i < pl; i++ {
		// Numerator (with truncation).
		// The original code was doing a floor division.
		// The truncation will have the same effect.
		n := numerator / d
		// Symmetric modulo.
		sm := symmetricModulo(n)
		// Add to the set of expansions.
		exp = append(exp, sm)
		if exp[i] < 0 {
			numerator = numerator + d*b
		}
		d *= b
	}
	return exp
}

func parseRational(rat float64, params *Parameters) int64 {
	// Absolute value of p.
	p := float64(-1 * params.MinPower())
	// Base.
	b := float64(params.Base())
	// Base to the power of minimal power: b^(|p|).
	bp := math.Pow(b, p)
	// Rational transformed.
	n := math.Trunc(rat * bp)
	return int64(n)
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

func rationalToFraction(r float64, params *Parameters) *big.Rat {
	// Absolute value of p.
	absP := math.Abs(float64(params.MinPower()))
	// Base to the power of the absolute value of p.
	bPowP := math.Pow(float64(params.Base()), absP)
	// Generate numerator by multiplying rational by b^{|p|} and taking the integer part.
	n := int64(math.Trunc(r * bPowP))
	// Generate denominator as b^{|p|}.
	d := int64(bPowP)
	// Generate fraction
	return big.NewRat(n, d)
}

// roundUp rounds a float up to p (minimal power) decimal places.
func roundUp(r float64, params *Parameters) float64 {
	// Base to the power of the absolute value of p.
	p := math.Pow(float64(params.Base()), float64(-1*params.MinPower()))
	// Digit.
	d := p * r
	return math.Ceil(d) / p
}
