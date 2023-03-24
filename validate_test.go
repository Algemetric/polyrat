package sim2d

import (
	"math/big"
	"testing"
)

func TestValidateDenominator(t *testing.T) {
	// Input values: fraction number (n), radix (r), expected modulo (em).
	d := big.NewInt(2401)
	// Create parameters (b, p, q, d).
	params, err := NewParameters(2, -8, 1, 16)
	if err != nil {
		t.Error(err)
	}
	// Validate.
	err = validateDenominator(d, params)
	if err == nil {
		t.Error("an error should be thrown when the denominator is not the base to the power of the absolute value of p")
	} else {
		if err.Error() != ErrDenominatorIsNotEqualToBToThePowerOfP.Error() {
			t.Error(ErrDenominatorIsNotEqualToBToThePowerOfP.Error())
		}
	}

	// Case for when denominator is 0.
	d = big.NewInt(0)
	// Parameters.
	params, err = NewParameters(10, -4, 1, 8)
	if err != nil {
		t.Error(err)
	}
	// Validate.
	err = validateDenominator(d, params)
	if err != ErrDenominatorIsZero {
		t.Error("a zero denominator should throw an error")
	}
}

func TestValidateNumerator(t *testing.T) {
	// b is EVEN.
	// Create parameters (b, p, q, d).
	params, err := NewParameters(10, -2, 3, 16)
	if err != nil {
		t.Error(err)
	}
	// Check that numerator -523187 is in the message space [-555555, 444444].
	n := big.NewInt(-523187)
	err = validateNumerator(n, params)
	if err != nil {
		t.Error("an error should not be thrown because the numerator is in the message space range")
	}
	// Check that numerator 455192 is NOT in the message space [-555555, 444444].
	n.SetInt64(455192)
	err = validateNumerator(n, params)
	if err == nil {
		t.Error("an error should be thrown because the numerator is not in the message space range")
	} else {
		if err.Error() != ErrNumeratorIsNotInTheMessageSpaceRange.Error() {
			t.Error(ErrNumeratorIsNotInTheMessageSpaceRange.Error())
		}
	}
	// b is ODD.
	// Create parameters (b, p, q, d).
	params, err = NewParameters(7, -2, 3, 16)
	if err != nil {
		t.Error(err)
	}
	// Check that numerator -58830 is not in the message space [,].
	n.SetInt64(-58830)
	err = validateNumerator(n, params)
	if err == nil {
		t.Error("an error should be thrown because the numerator is not in the message space range")
	} else {
		if err.Error() != ErrNumeratorIsNotInTheMessageSpaceRange.Error() {
			t.Error(ErrNumeratorIsNotInTheMessageSpaceRange.Error())
		}
	}
	// Check that numerator -58823 is in the message space [,].
	n.SetInt64(-58823)
	err = validateNumerator(n, params)
	if err != nil {
		t.Error("an error should not be thrown because the numerator is in the message space range")
	}
	// Check that numerator 58832 is not in the message space [,].
	n.SetInt64(58832)
	err = validateNumerator(n, params)
	if err == nil {
		t.Error("an error should be thrown because the numerator is not in the message space range")
	} else {
		if err.Error() != ErrNumeratorIsNotInTheMessageSpaceRange.Error() {
			t.Error(ErrNumeratorIsNotInTheMessageSpaceRange.Error())
		}
	}
	// Check that numerator 58773 is in the message space [,].
	n.SetInt64(58773)
	err = validateNumerator(n, params)
	if err != nil {
		t.Error("an error should not be thrown because the numerator is in the message space range")
	}
}

func TestValidateDegreeOfCode(t *testing.T) {
	// Input code.
	c := []int64{2, -3, 0, 0, -3, 0, 1}
	// Create parameters (b, p, q, d).
	params, err := NewParameters(7, -2, 3, 16)
	if err != nil {
		t.Error(err)
	}
	// Decoded fraction.
	err = validateDegreeOfCode(c, params)
	if err == nil {
		t.Error("an error should be thrown when the degree of the code is not a power of 2")
	} else {
		if err.Error() != ErrCodeDegreeIsNotAPowerOfTwo.Error() {
			t.Error(ErrCodeDegreeIsNotAPowerOfTwo.Error())
		}
	}
	// Check if the code received has the right degree.
	// Code should be of size 16, but is 32.
	c = []int64{5, 5, 0, -3, -3, 1, 0, 0, 0, 0, 0, 0, -5, -14, -26, -40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// Check code degree.
	err = validateDegreeOfCode(c, params)
	if err == nil {
		t.Error("an error should be thrown when the code has a different degree")
	} else {
		if err.Error() != ErrCodeDegreeIsDifferentFromDegree.Error() {
			t.Error(ErrCodeDegreeIsDifferentFromDegree.Error())
		}
	}
}

func TestValidateEncodingParameters(t *testing.T) {
	// Valid fraction.
	num := big.NewInt(-44979)
	den := big.NewInt(2401)
	// Create parameters (b, p, q, d).
	params, err := NewParameters(7, -4, 1, 8)
	if err != nil {
		t.Error(err)
	}
	// Validate.
	err = validateEncodingParameters(num, den, params)
	if err != nil {
		t.Error("parameters should be valid for encoding")
	}
}

func TestValidateDecodingParameters(t *testing.T) {
	// Create parameters (b, p, q, d).
	params, err := NewParameters(7, -4, 1, 8)
	if err != nil {
		t.Error(err)
	}
	// Expected code.
	c := []int64{2, -3, 0, 0, -3, 0, 1, -2}
	err = validateDecodingParameters(c, params)
	if err != nil {
		t.Error("parameters should be valid for decoding")
	}
}
