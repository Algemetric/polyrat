package sim2d

import (
	"strings"
	"testing"
)

// TestDecode tests the decoding of a polynomial into a rational.
func TestDecode(t *testing.T) {
	// Expected rational for code: 98123.45.
	er := 98123.45
	// Code.
	c := []int64{4, 2, 1, -2, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 5, 5}
	// Create parameters (p, q, d).
	params, err := NewParameters(-4, 11, 16)
	if err != nil {
		t.Error(err)
	}
	// Decoded rational.
	dr, err := Decode(c, params)
	if err != nil {
		t.Error(err)
	}
	if dr != er {
		t.Errorf("error decoding, expected %f but got %f", er, dr)
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
