package sim2d

import (
	"math/big"
	"testing"
)

func TestValidateModulo(t *testing.T) {
	// b >= 2.
	b := 1
	err := validateModulo(b)
	if err == nil {
		t.Error("an error should be thrown when b is less than 2")
	} else {
		if err.Error() != ErrBIsLessThan2.Error() {
			t.Error(ErrBIsLessThan2.Error())
		}
	}
}

func TestValidateSmallestPowerOfExpansion(t *testing.T) {
	// p < q.
	p, q := 1, 1
	err := validateSmallestPowerOfExpansion(p, q)
	if err == nil {
		t.Error("an error should be thrown when p is less than q")
	} else {
		if err.Error() != ErrPIsLessThanQ.Error() {
			t.Error(ErrPIsLessThanQ.Error())
		}
	}
	// p < 0.
	p, q = 0, 2
	err = validateSmallestPowerOfExpansion(p, q)
	if err == nil {
		t.Error("an error should be thrown when p is greater than or equal to 0")
	} else {
		if err.Error() != ErrPIsGreaterThanOrEqualToZero.Error() {
			t.Error(ErrPIsGreaterThanOrEqualToZero.Error())
		}
	}
}

func TestValidateGreatestPowerOfExpansion(t *testing.T) {
	// q > 0.
	q := 0
	err := validateGreatestPowerOfExpansion(q)
	if err == nil {
		t.Error("an error should be thrown when q is less than or equal to 0")
	} else {
		if err.Error() != ErrQIsLessThanOrEqualToZero.Error() {
			t.Error(ErrQIsLessThanOrEqualToZero.Error())
		}
	}
}

func TestValidateDegree(t *testing.T) {
	// d is a power of 2.
	p, q, d := -4, 1, 3
	err := validateDegree(p, q, d)
	if err == nil {
		t.Error("an error should be thrown when d is not a power of 2")
	} else {
		if err.Error() != ErrDIsNotAPowerOf2.Error() {
			t.Error(ErrDIsNotAPowerOf2.Error())
		}
	}
	// d >= 1.
	p, q, d = -4, 1, 0
	err = validateDegree(p, q, d)
	if err == nil {
		t.Error("an error should be thrown when d is less than 1")
	} else {
		if err.Error() != ErrDIsLessThanOne.Error() {
			t.Error(ErrDIsLessThanOne.Error())
		}
	}
	// d > q + |p|.
	p, q, d = 1, 7, 8
	err = validateDegree(p, q, d)
	if err == nil {
		t.Error("an error should be thrown when d is less than or equal to q plus the absolute value of p")
	} else {
		if err.Error() != ErrDIsLessThanOrEqualToQPlusP.Error() {
			t.Error(ErrDIsLessThanOrEqualToQPlusP.Error())
		}
	}
}

func TestValidateDenominator(t *testing.T) {
	// Input values: fraction number (n), radix (r), expected modulo (em).
	f := big.NewRat(-44979, 2401)
	b, p := 1, 8
	err := validateDenominator(f, b, p)
	if err == nil {
		t.Error("an error should be thrown when the denominator is not the base to the power of the absolute value of p")
	} else {
		if err.Error() != ErrDenominatorIsNotEqualToBToThePowerOfP.Error() {
			t.Error(ErrDenominatorIsNotEqualToBToThePowerOfP.Error())
		}
	}
}

func TestValidateNumerator(t *testing.T) {
	// Input values: number (n), radix (r).
	f := big.NewRat(-58825, 2401)
	b, p, q := 7, -4, 1
	err := validateNumerator(f, b, p, q)
	if err == nil {
		t.Error("an error should be thrown when the numerator is not in the message space range")
	} else {
		if err.Error() != ErrNumeratorIsNotInTheMessageSpaceRange.Error() {
			t.Error(ErrNumeratorIsNotInTheMessageSpaceRange.Error())
		}
	}
}

func TestValidateEncodingParameters(t *testing.T) {
	// Valid fraction.
	f := big.NewRat(-44979, 2401)
	// Valid parameters.
	b, p, q, d := 7, -4, 1, 8
	err := validateEncodingParameters(f, b, p, q, d)
	if err != nil {
		t.Error("parameters should be valid")
	}
}
