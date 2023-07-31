package polyrat

import (
	"strings"
	"testing"
)

// TestDecode tests the decoding of a polynomial into a rational.
func TestDecode(t *testing.T) {
	// Case for a special rational 83740034.866.
	er := 83740034.866
	// b = 10, p = -8, q = 12, d = 2048.
	// Create parameters (p, q, d).
	params, err := NewParameters(-8, 20, 2048)
	if err != nil {
		t.Error(err)
	}
	// Encode.
	c, err := Encode(er, params)
	if err != nil {
		t.Error(err)
	}
	// Decode.
	dr, err := Decode(c, params)
	if err != nil {
		t.Error(err)
	}
	// Check result.
	if dr != er {
		t.Errorf("error decoding, expected %f but got %f", er, dr)
	}

	// Expected rational for code: 98123.45.
	er = 98123.45
	// Code.
	c = []int64{4, 2, 1, -2, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 5, 5}
	// Create parameters (p, q, d).
	params, err = NewParameters(-4, 11, 16)
	if err != nil {
		t.Error(err)
	}
	// Decoded rational.
	dr, err = Decode(c, params)
	if err != nil {
		t.Error(err)
	}
	if dr != er {
		t.Errorf("error decoding, expected %f but got %f", er, dr)
	}

	// Case for different decimal numbers.
	// Encode every number to check the decoding and rounding.
	r := []float64{123.01, 123.12, 123.23, 123.34, 123.45, 123.56, 123.67, 123.78, 123.89, 123.90}
	// Create parameters (p, q, d).
	params, err = NewParameters(-2, 11, 16)
	if err != nil {
		t.Error(err)
	}
	// Encode.
	var cc [][]int64
	for i := 0; i < len(r); i++ {
		c, err := Encode(r[i], params)
		if err != nil {
			t.Error(err)
			break
		}
		cc = append(cc, c)
	}
	// Decode.
	for i := 0; i < len(cc); i++ {
		dr, err := Decode(cc[i], params)
		if err != nil {
			t.Error(err)
			break
		}
		if dr != r[i] {
			t.Errorf("error decoding, expected %f but got %f", r[i], dr)
		}
	}
}

func TestEvaluationPowers(t *testing.T) {
	// Create parameters (b, p, q, d).
	params, err := NewParameters(-4, 11, 16)
	if err != nil {
		t.Error(err)
	}
	// Expected results:
	er := []string{"1/10000", "1/1000", "1/100", "1/10", "1/1", "10/1", "100/1", "1000/1", "10000/1", "100000/1",
		"1000000/1", "10000000/1", "100000000/1", "1000000000/1", "10000000000/1", "100000000000/1"}
	ep := evaluationPowers(params)
	for i := 0; i < len(er); i++ {
		if strings.Compare(ep[i].String(), er[i]) != 0 {
			t.Errorf("expected %s value at position %d, but got %s", er[i], i, ep[i].String())
		}
	}
}
