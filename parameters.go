package polyrat

import "math"

// Parameters struct organizes the base, great power, least power
// and polynomial degree information given to the encoding and
// decoding functions.
type Parameters struct {
	base       int // Base (modulo).
	greatPower int // Great power.
	leastPower int // Least power.
	degree     int // Degree of the polynomial.
}

// NewParameters creates a struct that validates all the parameters
// used for encoding and decoding.
func NewParameters(leastPower, greatPower, degree int) (*Parameters, error) {
	// Setting up given parameters.
	params := new(Parameters)
	params.base = Base
	params.greatPower = greatPower
	params.leastPower = leastPower
	params.degree = degree
	// Validation of parameters.
	err := params.validate()
	if err != nil {
		return params, err
	}
	return params, nil
}

// Getter for base.
func (params *Parameters) Base() int {
	return params.base
}

// Getter for great power.
func (params *Parameters) GreatPower() int {
	return params.greatPower
}

// Getter for least power.
func (params *Parameters) LeastPower() int {
	return params.leastPower
}

// Getter for degree.
func (params *Parameters) Degree() int {
	return params.degree
}

// validateLeastPower validates criteria for the least power of expansion.
func (params *Parameters) validateLeastPower() error {
	// leastPower < greatPower.
	if params.leastPower >= params.greatPower {
		return ErrLeastPowerIsLessThanGreatPower
	}
	// leastPower < 0.
	if params.leastPower >= 0 {
		return ErrLeastPowerIsGreaterThanOrEqualToZero
	}
	return nil
}

// validateGreatPower validates criteria for the greatest power of expansion.
func (params *Parameters) validateGreatPower() error {
	// greatPower > 0.
	if params.greatPower <= 0 {
		return ErrGreatPowerIsLessThanOrEqualToZero
	}
	return nil
}

// validateDegree validates criteria for the polynomial degree.
func (params *Parameters) validateDegree() error {
	// degree >= 1.
	if params.degree < 1 {
		return ErrDegreeIsLessThanOne
	}
	// degree is a power of 2.
	if degreeIsNotAPowerOfTwo(params.degree) {
		return ErrDegreeIsNotAPowerOfTwo
	}
	// degree > greatPower + |leastPower|.
	absP := math.Abs(float64(params.leastPower))
	if params.degree <= (params.greatPower + int(absP)) {
		return ErrDegreeIsNotGreaterThanGreatPowerPlusLeastPower
	}
	return nil
}

// validate is a general function that checks all parameters.
func (params *Parameters) validate() error {
	// Error variable.
	var err error
	// Validades least power of expansion.
	err = params.validateLeastPower()
	if err != nil {
		return err
	}
	// Validades great power of expansion.
	err = params.validateGreatPower()
	if err != nil {
		return err
	}
	// Validates degree.
	err = params.validateDegree()
	if err != nil {
		return err
	}
	return nil
}
