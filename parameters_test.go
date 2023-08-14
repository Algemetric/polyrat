package polyrat

import (
	"testing"
)

func TestValidateLeastPower(t *testing.T) {
	// Check if an error is thrown when least power is less than great power.
	// Create parameters.
	_, err := NewParameters(1, 1, 0)
	// leastPower is valid if < greatPower.
	if err == nil {
		t.Error("an error should be thrown when least power is less than great power")
	} else {
		if err.Error() != ErrLeastPowerIsLessThanGreatPower.Error() {
			t.Error(ErrLeastPowerIsLessThanGreatPower.Error())
		}
	}

	// Check if an error is thrown when least power is greater than or equal to 0.
	_, err = NewParameters(0, 2, 0)
	if err == nil {
		t.Error("an error should be thrown when least power is greater than or equal to 0")
	} else {
		if err.Error() != ErrLeastPowerIsGreaterThanOrEqualToZero.Error() {
			t.Error(ErrLeastPowerIsGreaterThanOrEqualToZero.Error())
		}
	}
}

func TestValidateGreatPower(t *testing.T) {
	// Check if an error is thrown when great power is less than or equal to 0.
	// Create parameters.
	_, err := NewParameters(-1, 0, 0)
	// Great power is valid if > 0.
	if err == nil {
		t.Error("an error should be thrown when great power is less than or equal to 0")
	} else {
		if err.Error() != ErrGreatPowerIsLessThanOrEqualToZero.Error() {
			t.Error(ErrGreatPowerIsLessThanOrEqualToZero.Error())
		}
	}
}

func TestValidateDegree(t *testing.T) {
	// Check if an error is thrown when degree is not a power of 2.
	// Create parameters.
	_, err := NewParameters(-4, 1, 3)
	// Degree is valid if it is a power of 2.
	if err == nil {
		t.Error("an error should be thrown when degree is not a power of 2")
	} else {
		if err.Error() != ErrDegreeIsNotAPowerOfTwo.Error() {
			t.Error(ErrDegreeIsNotAPowerOfTwo.Error())
		}
	}

	// Check if an error is thrown when degree is less than 1.
	_, err = NewParameters(-4, 1, 0)
	if err == nil {
		t.Error("an error should be thrown when degree is less than 1")
	} else {
		if err.Error() != ErrDegreeIsLessThanOne.Error() {
			t.Error(ErrDegreeIsLessThanOne.Error())
		}
	}

	// Check if an error is thrown when degree is less than or equal to great power
	// plus the absolute value of the least power (degree <= greatPower + |leastPower|).
	_, err = NewParameters(-1, 9, 8)
	if err == nil {
		t.Error("an error should be thrown when degree is less than or equal to the great power plus the absolute value of the least power")
	} else {
		if err.Error() != ErrDegreeIsNotGreaterThanGreatPowerPlusLeastPower.Error() {
			t.Error(ErrDegreeIsNotGreaterThanGreatPowerPlusLeastPower.Error())
		}
	}
}
