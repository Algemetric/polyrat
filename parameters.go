package sim2d

import "math"

// Parameters struct organizes the base, high power, low power
// and polynomial degree information given to the encoding and
// decoding functions.
type Parameters struct {
	B int // b is the base (modulo).
	Q int // q is the higher power.
	P int // p is the lower power.
	D int // d is the degree of the polynomial.
}

// NewParameters creates a struct that validates all the parameters
// used for encoding and decoding.
func NewParameters(b, p, q, d int) (*Parameters, error) {
	// Setting up given parameters.
	// These inputs are set so that they can inform
	// the user about the parameters that caused an error.
	params := new(Parameters)
	params.B = b
	params.Q = q
	params.P = p
	params.D = d
	// Validation of parameters.
	err := params.validate()
	if err != nil {
		return params, err
	}
	return params, nil
}

// validateB validates criteria for the base.
func (params *Parameters) validateB() error {
	// b >= 2.
	if params.B < 2 {
		return ErrBIsLessThanTwo
	}
	return nil
}

// validateP validates criteria for the smallest power of expansion.
func (params *Parameters) validateP() error {
	// p < q.
	if params.P >= params.Q {
		return ErrPIsLessThanQ
	}
	// p < 0.
	if params.P >= 0 {
		return ErrPIsGreaterThanOrEqualToZero
	}
	return nil
}

// validateQ validates criteria for the greatest power of expansion.
func (params *Parameters) validateQ() error {
	// q > 0.
	if params.Q <= 0 {
		return ErrQIsLessThanOrEqualToZero
	}
	return nil
}

// validateD validates criteria for the polynomial degree.
func (params *Parameters) validateD() error {
	// d >= 1.
	if params.D < 1 {
		return ErrDIsLessThanOne
	}
	// d is a power of 2.
	// Log base 2 of d.
	floatD := float64(params.D)
	logD := math.Log2(floatD)
	// Integer part of log base 2 of d.
	intLog := math.Round(logD)
	// Recalculated d.
	d2 := math.Pow(2.0, intLog)
	if floatD != d2 {
		return ErrDIsNotAPowerOfTwo
	}
	// d > q + |p|.
	absP := math.Abs(float64(params.P))
	if params.D <= (params.Q + int(absP)) {
		return ErrDIsLessThanOrEqualToQPlusP
	}
	return nil
}

// validate is a general function that checks all parameters.
func (params *Parameters) validate() error {
	// Error variable.
	var err error
	// Validates base.
	err = params.validateB()
	if err != nil {
		return err
	}
	// Validades smallest power of expansion.
	err = params.validateP()
	if err != nil {
		return err
	}
	// Validades smallest power of expansion.
	err = params.validateQ()
	if err != nil {
		return err
	}
	// Validates degree.
	err = params.validateD()
	if err != nil {
		return err
	}
	return nil
}
