package sim2d

import (
	"testing"
)

func TestValidateB(t *testing.T) {
	// Create parameters (b, p, q, d).
	_, err := NewParameters(1, 0, 0, 0)
	// b is valid if >= 2.
	if err == nil {
		t.Error("an error should be thrown when b is less than 2")
	} else {
		if err.Error() != ErrBIsLessThanTwo.Error() {
			t.Error(ErrBIsLessThanTwo.Error())
		}
	}
}

func TestValidateP(t *testing.T) {
	// Create parameters (b, p, q, d).
	_, err := NewParameters(2, 1, 1, 0)
	// p is valid if < q.
	if err == nil {
		t.Error("an error should be thrown when p is less than q")
	} else {
		if err.Error() != ErrPIsLessThanQ.Error() {
			t.Error(ErrPIsLessThanQ.Error())
		}
	}
	// p is valid if < 0.
	_, err = NewParameters(2, 0, 2, 0)
	if err == nil {
		t.Error("an error should be thrown when p is greater than or equal to 0")
	} else {
		if err.Error() != ErrPIsGreaterThanOrEqualToZero.Error() {
			t.Error(ErrPIsGreaterThanOrEqualToZero.Error())
		}
	}
}

func TestValidateQ(t *testing.T) {
	// Create parameters (b, p, q, d).
	_, err := NewParameters(2, -1, 0, 0)
	// q is valid if > 0.
	if err == nil {
		t.Error("an error should be thrown when q is less than or equal to 0")
	} else {
		if err.Error() != ErrQIsLessThanOrEqualToZero.Error() {
			t.Error(ErrQIsLessThanOrEqualToZero.Error())
		}
	}
}

func TestValidateD(t *testing.T) {
	// Create parameters (b, p, q, d).
	_, err := NewParameters(2, -4, 1, 3)
	// d is valid if it is a power of 2.
	if err == nil {
		t.Error("an error should be thrown when d is not a power of 2")
	} else {
		if err.Error() != ErrDIsNotAPowerOfTwo.Error() {
			t.Error(ErrDIsNotAPowerOfTwo.Error())
		}
	}
	// d is valid if it is >= 1.
	_, err = NewParameters(2, -4, 1, 0)
	if err == nil {
		t.Error("an error should be thrown when d is less than 1")
	} else {
		if err.Error() != ErrDIsLessThanOne.Error() {
			t.Error(ErrDIsLessThanOne.Error())
		}
	}
	// d is valid if > q + |p|.
	_, err = NewParameters(2, -1, 9, 8)
	if err == nil {
		t.Error("an error should be thrown when d is less than or equal to q plus the absolute value of p")
	} else {
		if err.Error() != ErrDIsLessThanOrEqualToQPlusP.Error() {
			t.Error(ErrDIsLessThanOrEqualToQPlusP.Error())
		}
	}
}
