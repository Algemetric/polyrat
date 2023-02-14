package sim2dcodec

import (
	"math/big"
	"strings"
	"testing"
)

// TestDecode tests the decoding of a polynomial into a fraction.
func TestDecode(t *testing.T) {
	t.Skip()
	f := big.NewRat(-44979, 2401)
	b, p, q, d := 7, -4, 1, 8
	c, err := Encode(f, b, p, q, d)
	if err != nil {
		t.Error(err)
	}
	rf := Decode(c, b, p, q)
	if strings.Compare(f.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", f.String(), rf.String())
	}
	// Case for when the fraction is still reducible and can cause errors.
	// The reduced form of 1460326978/1331 is 12068818/11.
	f = big.NewRat(1460326978, 1331)
	b, p, q, d = 11, -3, 6, 16
	c, err = Encode(f, b, p, q, d)
	if err != nil {
		t.Error(err)
	}
	rf = Decode(c, b, p, q)
	if strings.Compare(f.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", f.String(), rf.String())
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
