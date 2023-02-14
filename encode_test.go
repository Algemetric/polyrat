package sim2dcodec

import (
	"math/big"
	"testing"
)

// TestEncode tests the encoding of a fractional number into a polynomial.
func TestEncode(t *testing.T) {
	t.Skip()
	f := big.NewRat(-44979, 2401)
	b, p, q, d := 7, -4, 1, 8
	// Calculate code.
	c, err := Encode(f, b, p, q, d)
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
	f = big.NewRat(1460326978, 1331)
	b, p, q, d = 11, -3, 6, 16
	c, err = Encode(f, b, p, q, d)
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
}
