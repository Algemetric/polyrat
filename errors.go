package polyrat

import "errors"

var (
	ErrDegreeIsNotGreaterThanGreatPowerPlusLeastPower = errors.New("degree should be greater than the great power plus the absolute value of the least power")
	ErrDegreeIsLessThanOne                            = errors.New("degree should be greater than or equal to 1")
	ErrDegreeIsNotAPowerOfTwo                         = errors.New("degree should be a power of 2")
	ErrLeastPowerIsLessThanGreatPower                 = errors.New("the least power should be less than the great power")
	ErrLeastPowerIsGreaterThanOrEqualToZero           = errors.New("the least power should be less than 0")
	ErrGreatPowerIsLessThanOrEqualToZero              = errors.New("great power should be greater than 0")
	ErrNumeratorIsNotInTheMessageSpace                = errors.New("numerator should be inside the message space range")
	ErrCodeDegreeIsNotAPowerOfTwo                     = errors.New("code degree should be a power of 2")
	ErrCodeDegreeIsDifferentFromDegree                = errors.New("code degree is different from the acceptable degree")
)
