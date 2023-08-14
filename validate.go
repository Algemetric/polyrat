package polyrat

import (
	"math"
)

// inputIsInvalid checks if the number given to the function is in the input space.
// We just consider the case when the parity of the base is even.
func inputIsInvalid(input int64, params *Parameters) bool {
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

func validateDegreeOfCode(code []int64, params *Parameters) error {
	// Code degree.
	cd := len(code)
	// Check if code degree is a power of 2.
	// Log base 2 of code degree.
	floatD := float64(cd)
	logD := math.Log2(floatD)
	// Integer part of log base 2 of code degree.
	intLog := math.Round(logD)
	// Code degree recalculated.
	cdr := math.Pow(2.0, intLog)
	// Check if code degree is a power of 2.
	if floatD != cdr {
		return ErrCodeDegreeIsNotAPowerOfTwo
	}
	// Check if code degree is the same as the given degree.
	if params.Degree() != cd {
		return ErrCodeDegreeIsDifferentFromDegree
	}
	return nil
}

func validateDecodingParameters(code []int64, params *Parameters) error {
	// Validate if degree of code is a power of 2.
	err := validateDegreeOfCode(code, params)
	if err != nil {
		return err
	}
	return nil
}
