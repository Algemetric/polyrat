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
	b, p, q, d := 7, -4, 1, 8
	// Calculate code.
	c, err := Encode(num, den, b, p, q, d)
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
	b, p, q, d = 11, -3, 6, 16
	c, err = Encode(num, den, b, p, q, d)
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
	b, p, q, d = 10, -4, 10, 16
	c, err = Encode(num, den, b, p, q, d)
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
	// Parameters.
	b, p, q, d = 10, -4, 8, 16
	c, err = Encode(num, den, b, p, q, d)
	if err != nil {
		t.Error(err)
	}
	// Expected fraction.
	ef := big.NewRat(670592745, 10000)
	// Decoded fraction.
	rf, err := Decode(c, b, p, q)
	if err != nil {
		t.Error(err)
	}
	if strings.Compare(ef.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", ef.String(), rf.String())
	}

	// Case for when denominator is 0.
	num = big.NewInt(1)
	den = big.NewInt(0)
	// Parameters.
	b, p, q, d = 10, -4, 8, 16
	c, err = Encode(num, den, b, p, q, d)
	if err != ErrDenominatorIsZero {
		t.Error("denominator cannot be zero")
	}
}
