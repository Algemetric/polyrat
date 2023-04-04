package sim2d

import "math"

// Parameters struct organizes the base, high power, low power
// and polynomial degree information given to the encoding and
// decoding functions.
type Parameters struct {
	b int // b is the base (modulo).
	q int // q is the higher power.
	p int // p is the lower power.
	d int // d is the degree of the polynomial.
}

// NewParameters creates a struct that validates all the parameters
// used for encoding and decoding.
func NewParameters(p, q, d int) (*Parameters, error) {
	// Setting up given parameters.
	// These inputs are set so that they can inform
	// the user about the parameters that caused an error.
	params := new(Parameters)
	params.b = Base
	params.q = q
	params.p = p
	params.d = d
	// Validation of parameters.
	err := params.validate()
	if err != nil {
		return params, err
	}
	return params, nil
}

// Getter for base.
func (params *Parameters) Base() int {
	return params.b
}

// Getter for higher power.
func (params *Parameters) MaxPower() int {
	return params.q
}

// Getter for lower power.
func (params *Parameters) MinPower() int {
	return params.p
}

// Getter for degree.
func (params *Parameters) Degree() int {
	return params.d
}

// validateP validates criteria for the smallest power of expansion.
func (params *Parameters) validateP() error {
	// p < q.
	if params.p >= params.q {
		return ErrPIsLessThanQ
	}
	// p < 0.
	if params.p >= 0 {
		return ErrPIsGreaterThanOrEqualToZero
	}
	return nil
}

// validateQ validates criteria for the greatest power of expansion.
func (params *Parameters) validateQ() error {
	// q > 0.
	if params.q <= 0 {
		return ErrQIsLessThanOrEqualToZero
	}
	return nil
}

// validateD validates criteria for the polynomial degree.
func (params *Parameters) validateD() error {
	// d >= 1.
	if params.d < 1 {
		return ErrDIsLessThanOne
	}
	// d is a power of 2.
	// Log base 2 of d.
	floatD := float64(params.d)
	logD := math.Log2(floatD)
	// Integer part of log base 2 of d.
	intLog := math.Round(logD)
	// Recalculated d.
	d2 := math.Pow(2.0, intLog)
	if floatD != d2 {
		return ErrDIsNotAPowerOfTwo
	}
	// d > q + |p|.
	absP := math.Abs(float64(params.p))
	if params.d <= (params.q + int(absP)) {
		return ErrDIsLessThanOrEqualToQPlusP
	}
	return nil
}

// validate is a general function that checks all parameters.
func (params *Parameters) validate() error {
	// Error variable.
	var err error
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
