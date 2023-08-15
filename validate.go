package polyrat

import (
	"math"
)

// inputNotInTheMessageSpace checks if the given number is in the message space.
// We just consider the case when the parity of the base is even.
func inputNotInTheMessageSpace(input int64, params *Parameters) bool {
	// Define a common component of all bounds: b^(gp-lp+1) - 1.
	b, gp, lp := float64(params.Base()), params.GreatPower(), params.LeastPower()
	// Exponent: e = gp-lp+1.
	e := float64(gp - lp + 1)
	// b^(gp-lp+1) - 1
	bp := math.Pow(b, e) - 1
	// We define the lower and upper bounds by defining the equations in separated parts.
	// Lower bound: -b/2 x (b^(gp-lp+1) - 1) / (b-1).
	lb := -b / 2 * bp / (b - 1)
	// Upper bound: (b/2 - 1) x ((b^(gp-lp+1) - 1) / (b-1)).
	ub := (b/2 - 1) * bp / (b - 1)
	// Check if number is less than lower bound or greater than upper bound.
	return input < int64(lb) || int64(ub) < input
}

func degreeIsNotAPowerOfTwo(d int) bool {
	// Code degree.
	cd := 1
	cd <<= int(math.Log2(float64(d)))
	// Return if degree is a power of 2.
	return d != cd
}

func validateCodeDegree(code []int64, params *Parameters) error {
	// Code length.
	cl := len(code)
	// Check if code degree is a power of 2.
	if degreeIsNotAPowerOfTwo(cl) {
		return ErrCodeDegreeIsNotAPowerOfTwo
	}
	// Check if code degree is the same as the given degree.
	if params.Degree() != cl {
		return ErrCodeDegreeIsDifferentFromDegree
	}
	return nil
}
