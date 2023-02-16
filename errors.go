package sim2d

import "errors"

var (
	ErrDIsLessThanOrEqualToQPlusP  = errors.New("degree should be greater than q plus the absolute value of p")
	ErrDIsLessThanOne              = errors.New("degree should be greater than or equal to 1")
	ErrDIsNotAPowerOf2             = errors.New("degree should be a power of 2")
	ErrPIsLessThanQ                = errors.New("p should be less than q")
	ErrPIsGreaterThanOrEqualToZero = errors.New("p should be less than 0")
	ErrQIsLessThanOrEqualToZero    = errors.New("q is greater than 0")
	ErrBIsLessThan2                = errors.New("base should be greater than or equal to 2")
)
