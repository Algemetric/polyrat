package sim2d

import "errors"

// b is the base.
// q is higher power.
// p is the lower power.
// d is the degree.

var (
	ErrDIsLessThanOrEqualToQPlusP           = errors.New("degree should be greater than the higher power plus the absolute value of the lower power")
	ErrDIsLessThanOne                       = errors.New("degree should be greater than or equal to 1")
	ErrDIsNotAPowerOfTwo                    = errors.New("degree should be a power of 2")
	ErrPIsLessThanQ                         = errors.New("the lower power should be less than the higher power")
	ErrPIsGreaterThanOrEqualToZero          = errors.New("the lower power should be less than 0")
	ErrQIsLessThanOrEqualToZero             = errors.New("higher power should be greater than 0")
	ErrNumeratorIsNotInTheMessageSpaceRange = errors.New("numerator should be inside the message space range")
	ErrCodeDegreeIsNotAPowerOfTwo           = errors.New("code degree should be a power of 2")
	ErrCodeDegreeIsDifferentFromDegree      = errors.New("code degree is different from the acceptable degree")
)
