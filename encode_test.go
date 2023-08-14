package polyrat

import (
	"testing"
)

// TestEncode tests the encoding of a rational number into a polynomial.
func TestEncode(t *testing.T) {
	// Rational number 98123.45.
	r := 98123.45
	// Create parameters.
	params, err := NewParameters(-4, 11, 16)
	if err != nil {
		t.Error(err)
	}
	c, err := Encode(r, params)
	if err != nil {
		t.Error(err)
	}
	// Expected code.
	ec := []int64{4, 2, 1, -2, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 5, 5}
	// Check length of code.
	if len(c) != len(ec) {
		t.Errorf("expected code has %d elements but got %d", len(ec), len(c))
	}
	// Check code values.
	for i := 0; i < len(ec); i++ {
		if ec[i] != c[i] {
			t.Errorf("expected code value %d at position %d but got %d", ec[i], i, c[i])
		}
	}

	// Check that numerator -523187 does not raise an error.
	r = -5231.87
	// Create parameters.
	params, err = NewParameters(-2, 3, 16)
	if err != nil {
		t.Error(err)
	}
	// Encode.
	_, err = Encode(r, params)
	if err != nil {
		t.Errorf("given rational should not raise an error, but got: %s", err.Error())
	}

	// Check that numerator 455192 raises an error.
	r = 4551.92
	// Encode.
	_, err = Encode(r, params)
	if err == nil {
		t.Errorf("given rational should raise an error")
	}
}
