package sim2dcodec

import (
	"math/big"
	"strings"
	"testing"
)

// TestDecode tests the decoding of a polynomial into a fraction.
func TestDecode(t *testing.T) {
	// Expected fraction.
	ef := big.NewRat(-44979, 2401)
	// Parameters.
	b, p, q := 7, -4, 1
	// Expected code.
	c := []int64{2, -3, 0, 0, -3, 0, 1, -2}
	// Decoded fraction.
	rf := Decode(c, b, p, q)
	if strings.Compare(ef.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", ef.String(), rf.String())
	}

	// Case for when the fraction is still reducible and can cause errors.
	// The reduced form of 1460326978/1331 is 12068818/11.
	// Expected fraction.
	ef = big.NewRat(1460326978, 1331)
	// Parameters.
	b, p, q = 11, -3, 6
	// Expected code.
	c = []int64{3, 5, 3, -1, -2, -4, 1, 0, 0, 0, 0, 0, 0, 0, 0, -3}
	// Decoded fraction.
	rf = Decode(c, b, p, q)
	if strings.Compare(ef.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", ef.String(), rf.String())
	}

	// Case for a larger fraction.
	// [1, 4, -1, 4, 3, 3, 0, 4, 3, 2, 0, 0, 3, 3, 1, -3] results in 23403339412867/10000.
	// Expected fraction.
	ef = big.NewRat(23403339412867, 10000)
	// Parameters.
	b, p, q = 10, -4, 10
	// Expected code.
	c = []int64{1, 4, -1, 4, 3, 3, 0, 4, 3, 2, 0, 0, 3, 3, 1, -3}
	// Decoded fraction.
	rf = Decode(c, b, p, q)
	if strings.Compare(ef.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", ef.String(), rf.String())
	}
}

func TestEvaluationPowers(t *testing.T) {
	b, p, q := 7, -4, 1
	// Expected results: (1/2401, 1/343, 1/49, 1/7, 1, 7).
	ep := evaluationPowers(b, q, p)
	if strings.Compare(ep[0].String(), "1/2401") != 0 ||
		strings.Compare(ep[1].String(), "1/343") != 0 ||
		strings.Compare(ep[2].String(), "1/49") != 0 ||
		strings.Compare(ep[3].String(), "1/7") != 0 ||
		strings.Compare(ep[4].String(), "1/1") != 0 ||
		strings.Compare(ep[5].String(), "7/1") != 0 {
		t.Error("error calculating evaluation powers")
	}
}
