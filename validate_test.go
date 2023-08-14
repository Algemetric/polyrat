package polyrat

import (
	"testing"
)

func TestValidateInput(t *testing.T) {
	// Check range when base is EVEN.
	// In our case base will always be even since we defined base 10 as a constant.
	// Create parameters.
	params, err := NewParameters(-2, 3, 16)
	if err != nil {
		t.Error(err)
	}
	// Check that input -523187 is in the message space [-555555, 444444].
	n := int64(-523187)
	if inputIsInvalid(n, params) {
		t.Error("an error should not be thrown because the numerator is in the message space range")
	}
	// Check that input 455192 is NOT in the message space [-555555, 444444].
	n = 455192
	if !inputIsInvalid(n, params) {
		t.Error("an error should be thrown because the numerator is not in the message space range")
	}
}

func TestValidateDegreeOfCode(t *testing.T) {
	// Check if an error is thrown when the degree of the code is not a power of 2.
	// Input code.
	c := []int64{2, -3, 0, 0, -3, 0, 1}
	// Create parameters.
	params, err := NewParameters(-2, 3, 16)
	if err != nil {
		t.Error(err)
	}
	// Decoded fraction.
	err = validateCodeDegree(c, params)
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
	err = validateCodeDegree(c, params)
	if err == nil {
		t.Error("an error should be thrown when the code has a different degree")
	} else {
		if err.Error() != ErrCodeDegreeIsDifferentFromDegree.Error() {
			t.Error(ErrCodeDegreeIsDifferentFromDegree.Error())
		}
	}
}

func TestValidateDecodingParameters(t *testing.T) {
	// Check that parameters are valid for decoding.
	// Create parameters.
	params, err := NewParameters(-4, 1, 8)
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
