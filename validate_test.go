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
	d := big.NewInt(2401)
	b, p := 1, -8
	err := validateDenominator(d, b, p)
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
	b, p = 10, -4
	err = validateDenominator(d, b, p)
	if err != ErrDenominatorIsZero {
		t.Error("a zero denominator should throw an error")
	}
}

func TestValidateNumerator(t *testing.T) {
	// b = 10, q = 3, p = -2.
	// b is EVEN.
	b, q, p := 10, 3, -2
	// Check that numerator -523187 is in the message space [-555555, 444444].
	n := big.NewInt(-523187)
	err := validateNumerator(n, b, p, q)
	if err != nil {
		t.Error("an error should not be thrown because the numerator is in the message space range")
	}
	// Check that numerator 455192 is NOT in the message space [-555555, 444444].
	n.SetInt64(455192)
	err = validateNumerator(n, b, p, q)
	if err == nil {
		t.Error("an error should be thrown because the numerator is not in the message space range")
	} else {
		if err.Error() != ErrNumeratorIsNotInTheMessageSpaceRange.Error() {
			t.Error(ErrNumeratorIsNotInTheMessageSpaceRange.Error())
		}
	}
	// b = 7, q = 3, p = -2.
	// b is ODD.
	b, q, p = 7, 3, -2
	// Check that numerator -58830 is not in the message space [,].
	n.SetInt64(-58830)
	err = validateNumerator(n, b, p, q)
	if err == nil {
		t.Error("an error should be thrown because the numerator is not in the message space range")
	} else {
		if err.Error() != ErrNumeratorIsNotInTheMessageSpaceRange.Error() {
			t.Error(ErrNumeratorIsNotInTheMessageSpaceRange.Error())
		}
	}
	// Check that numerator -58823 is in the message space [,].
	n.SetInt64(-58823)
	err = validateNumerator(n, b, p, q)
	if err != nil {
		t.Error("an error should not be thrown because the numerator is in the message space range")
	}
	// Check that numerator 58832 is not in the message space [,].
	n.SetInt64(58832)
	err = validateNumerator(n, b, p, q)
	if err == nil {
		t.Error("an error should be thrown because the numerator is not in the message space range")
	} else {
		if err.Error() != ErrNumeratorIsNotInTheMessageSpaceRange.Error() {
			t.Error(ErrNumeratorIsNotInTheMessageSpaceRange.Error())
		}
	}
	// Check that numerator 58773 is in the message space [,].
	n.SetInt64(58773)
	err = validateNumerator(n, b, p, q)
	if err != nil {
		t.Error("an error should not be thrown because the numerator is in the message space range")
	}
}

func TestValidateDegreeOfCode(t *testing.T) {
	// Input code.
	c := []int64{2, -3, 0, 0, -3, 0, 1}
	// Decoded fraction.
	err := validateDegreeOfCode(c)
	if err == nil {
		t.Error("an error should be thrown when the degree of the code is not a power of 2")
	} else {
		if err.Error() != ErrCodeDegreeIsNotAPowerOf2.Error() {
			t.Error(ErrCodeDegreeIsNotAPowerOf2.Error())
		}
	}
}

func TestValidateEncodingParameters(t *testing.T) {
	// Valid fraction.
	num := big.NewInt(-44979)
	den := big.NewInt(2401)
	// Valid parameters.
	b, p, q, d := 7, -4, 1, 8
	err := validateEncodingParameters(num, den, b, p, q, d)
	if err != nil {
		t.Error("parameters should be valid for encoding")
	}
}

func TestValidateDecodingParameters(t *testing.T) {
	// Parameters.
	b, p, q := 7, -4, 1
	// Expected code.
	c := []int64{2, -3, 0, 0, -3, 0, 1, -2}
	err := validateDecodingParameters(c, b, p, q)
	if err != nil {
		t.Error("parameters should be valid for decoding")
	}
}
