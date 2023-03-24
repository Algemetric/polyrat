package sim2d

import (
	"math/big"
	"strings"
	"testing"
)

// TestEncode tests the encoding of a fractional number into a polynomial.
func TestEncode(t *testing.T) {
	num := big.NewInt(-44979)
	den := big.NewInt(2401)
	// Create parameters (b, p, q, d).
	params, err := NewParameters(7, -4, 1, 8)
	if err != nil {
		t.Error(err)
	}
	// Calculate code.
	c, err := Encode(num, den, params)
	if err != nil {
		t.Error(err)
	}
	// Expected code.
	ec := []int64{2, -3, 0, 0, -3, 0, 1, -2}
	// Check results.
	for i := 0; i < len(c); i++ {
		if c[i] != ec[i] {
			t.Errorf("expected %d at position %d but got %d", ec[i], i, c[i])
		}
	}

	// Case for when the fraction is reducible and can cause errors.
	// The reduced form of 1460326978/1331 is 12068818/11.
	num = big.NewInt(1460326978)
	den = big.NewInt(1331)
	// Create parameters (b, p, q, d).
	params, err = NewParameters(11, -3, 6, 16)
	if err != nil {
		t.Error(err)
	}
	c, err = Encode(num, den, params)
	if err != nil {
		t.Error(err)
	}
	// Expected code.
	ec = []int64{3, 5, 3, -1, -2, -4, 1, 0, 0, 0, 0, 0, 0, 0, 0, -3}
	// Check results.
	for i := 0; i < len(c); i++ {
		if c[i] != ec[i] {
			t.Errorf("expected %d at position %d but got %d", ec[i], i, c[i])
		}
	}

	// Case for a larger fraction.
	// 23403339412867/10000 results in [1, 4, -1, 4, 3, 3, 0, 4, 3, 2, 0, 0, 3, 3, 1, -3].
	num = big.NewInt(23403339412867)
	den = big.NewInt(10000)
	// Create parameters (b, p, q, d).
	params, err = NewParameters(10, -4, 10, 16)
	if err != nil {
		t.Error(err)
	}
	c, err = Encode(num, den, params)
	if err != nil {
		t.Error(err)
	}
	// Expected code.
	ec = []int64{1, 4, -1, 4, 3, 3, 0, 4, 3, 2, 0, 0, 3, 3, 1, -3}
	// Check results.
	for i := 0; i < len(c); i++ {
		if c[i] != ec[i] {
			t.Errorf("expected %d at position %d but got %d", ec[i], i, c[i])
		}
	}

	// Case for when the value to be encoded is 67059.2745.
	num = big.NewInt(670592745)
	den = big.NewInt(10000)
	// Create parameters (b, p, q, d).
	params, err = NewParameters(10, -4, 8, 16)
	if err != nil {
		t.Error(err)
	}
	c, err = Encode(num, den, params)
	if err != nil {
		t.Error(err)
	}
	// Expected fraction.
	ef := big.NewRat(670592745, 10000)
	// Decoded fraction.
	rf, err := Decode(c, params)
	if err != nil {
		t.Error(err)
	}
	if strings.Compare(ef.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", ef.String(), rf.String())
	}

	// Case for when denominator is 0.
	// Parameters are still 10, -4, 8, 16 (b, p, q, d).
	num = big.NewInt(1)
	den = big.NewInt(0)
	c, err = Encode(num, den, params)
	if err != ErrDenominatorIsZero {
		t.Error("denominator cannot be zero")
	}
}
