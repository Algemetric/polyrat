package polyrat

import (
	"math"
)

// inputIsInvalid checks if the number given to the function is in the input space.
// We just consider the case when the parity of the base is even.
func inputIsInvalid(input int64, params *Parameters) bool {
	// Define a common component of all bounds: b^(q-p+1) - 1.
	b, q, p := float64(params.Base()), params.MaxPower(), params.MinPower()
	// Exponent: e = q-p+1.
	e := float64(q - p + 1)
	// b^(q-p+1) - 1
	bp := math.Pow(b, e) - 1
	// We define the lower and upper bounds by defining the equations in separated parts.
	// Lower bound: -b/2 x (b^(q-p+1) - 1) / (b-1).
	lb := ((-b / 2) * bp) / (b - 1)
	// Upper bound: (b/2 - 1) x ((b^(q-p+1) - 1) / (b-1)).
	ub := (((b / 2) - 1) * bp) / (b - 1)
	// Check if number is less than lower bound or greater than upper bound.
	return float64(input) < lb || ub < float64(input)
}

func validateDegreeOfCode(code []int64, params *Parameters) error {
	// Code degree.
	cd := len(code)
	// Check if d is a power of 2.
	// Log base 2 of d.
	floatD := float64(cd)
	logD := math.Log2(floatD)
	// Integer part of log base 2 of d.
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
