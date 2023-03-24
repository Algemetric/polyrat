package sim2d

import (
	"math/big"
	"strings"
	"testing"
)

// TestDecode tests the decoding of a polynomial into a fraction.
func TestDecode(t *testing.T) {
	// Expected fraction.
	ef := big.NewRat(-44979, 2401)
	// Create parameters (b, p, q, d).
	params, err := NewParameters(7, -4, 1, 8)
	if err != nil {
		t.Error(err)
	}
	// Expected code.
	c := []int64{2, -3, 0, 0, -3, 0, 1, -2}
	// Decoded fraction.
	rf, err := Decode(c, params)
	if err != nil {
		t.Error(err)
	}
	if strings.Compare(ef.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", ef.String(), rf.String())
	}

	// Case for when the fraction is still reducible and can cause errors.
	// The reduced form of 1460326978/1331 is 12068818/11.
	// Expected fraction.
	ef = big.NewRat(1460326978, 1331)
	// Create parameters (b, p, q, d).
	params, err = NewParameters(11, -3, 6, 16)
	if err != nil {
		t.Error(err)
	}
	// Expected code.
	c = []int64{3, 5, 3, -1, -2, -4, 1, 0, 0, 0, 0, 0, 0, 0, 0, -3}
	// Decoded fraction.
	rf, err = Decode(c, params)
	if err != nil {
		t.Error(err)
	}
	if strings.Compare(ef.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", ef.String(), rf.String())
	}

	// Case for a larger fraction.
	// [1, 4, -1, 4, 3, 3, 0, 4, 3, 2, 0, 0, 3, 3, 1, -3] results in 23403339412867/10000.
	// Expected fraction.
	ef = big.NewRat(23403339412867, 10000)
	// Create parameters (b, p, q, d).
	params, err = NewParameters(10, -4, 10, 16)
	if err != nil {
		t.Error(err)
	}
	// Expected code.
	c = []int64{1, 4, -1, 4, 3, 3, 0, 4, 3, 2, 0, 0, 3, 3, 1, -3}
	// Decoded fraction.
	rf, err = Decode(c, params)
	if err != nil {
		t.Error(err)
	}
	if strings.Compare(ef.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", ef.String(), rf.String())
	}

	// Case for when the value to be decoded is 10717.02.
	c = []int64{-2, -8, 8, -10, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, -2, 10}
	// Create parameters (b, p, q, d).
	params, err = NewParameters(10, -4, 8, 16)
	if err != nil {
		t.Error(err)
	}
	// Expected fraction.
	ef = big.NewRat(107170200, 10000)
	// Decoded fraction.
	rf, err = Decode(c, params)
	if err != nil {
		t.Error(err)
	}
	if strings.Compare(ef.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", ef.String(), rf.String())
	}

	// Case for when the value to be decoded is 67059.2745.
	c = []int64{5, 5, 0, -3, -3, 1, 0, 0, 0, 0, 0, 0, -5, -14, -26, -40}
	// Create parameters (b, p, q, d).
	params, err = NewParameters(10, -4, 8, 16)
	if err != nil {
		t.Error(err)
	}
	// Expected fraction.
	ef = big.NewRat(670592745, 10000)
	// Decoded fraction.
	rf, err = Decode(c, params)
	if err != nil {
		t.Error(err)
	}
	if strings.Compare(ef.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", ef.String(), rf.String())
	}
}

func TestEvaluationPowers(t *testing.T) {
	// Create parameters (b, p, q, d).
	params, err := NewParameters(7, -4, 1, 16)
	if err != nil {
		t.Error(err)
	}
	// Expected results: (1/2401, 1/343, 1/49, 1/7, 1, 7).
	ep := evaluationPowers(params)
	if strings.Compare(ep[0].String(), "1/2401") != 0 ||
		strings.Compare(ep[1].String(), "1/343") != 0 ||
		strings.Compare(ep[2].String(), "1/49") != 0 ||
		strings.Compare(ep[3].String(), "1/7") != 0 ||
		strings.Compare(ep[4].String(), "1/1") != 0 ||
		strings.Compare(ep[5].String(), "7/1") != 0 {
		t.Error("error calculating evaluation powers")
	}
}
