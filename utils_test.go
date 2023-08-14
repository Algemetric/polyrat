package polyrat

import (
	"math/big"
	"testing"
)

func TestSymmetricModulo(t *testing.T) {
	// Create parameters (p, q, d).
	p, err := NewParameters(-4, 11, 16)
	if err != nil {
		t.Error(err)
	}
	// Expected modules.
	em := []int64{0, 0, -5, 4, 3, 2, 1, -2, -1, 0, 0, 0, 0, 0, 0, 0}
	// Polynomial length.
	pl := polynomialLength(p)
	// Rational number and first denominator of the progression (d^0=1, d^1=10, d^2=100, ...).
	r := 981234500.0
	d := 1.0

	for i := 0; i < pl; i++ {
		n := int64(r / d)
		m := symmetricModulo(n, p)
		if em[i] != m {
			t.Errorf("expected %d but got %d", em[i], m)
		}
		d *= float64(p.Base())
	}
}

func TestExpansion(t *testing.T) {
	// Create parameters (p, q, d).
	params, err := NewParameters(-4, 11, 16)
	if err != nil {
		t.Error(err)
	}
	// Rational input value.
	// Numerator after being separated from the fraction.
	n := int64(981234500)
	// Calculate expansion.
	e := expansion(n, params)
	// Expected expansion.
	ee := []int64{0, 0, -5, -5, 4, 2, 1, -2, 0, 1, 0, 0, 0, 0, 0, 0}
	// Check if expansions have the same size.
	if len(ee) != len(e) {
		t.Errorf("expected expansion has %d elements but got %d", len(ee), len(e))
	}
	// Check if calculated expansion matches the expected values.
	for i := 0; i < len(ee); i++ {
		if e[i] != ee[i] {
			t.Errorf("expected expansion of %v but got %v", ee, e)
			break
		}
	}
}

func TestRationalToFraction(t *testing.T) {
	// Rational input value.
	r := 1097165.2727272727
	// Create parameters (p, q, d).
	params, err := NewParameters(-3, 1, 16)
	if err != nil {
		t.Error(err)
	}
	// Generate fraction.
	f := rationalToFraction(r, params)
	// Expected fraction.
	ef := big.NewRat(1097165272, 1000)
	if f.Cmp(ef) != 0 {
		t.Errorf("expected fraction %s but got %s", ef.String(), f.String())
	}
}

func TestRoundUp(t *testing.T) {
	// Rational value.
	r := 1097165.2727272727
	// Expected rounded rational.
	er := 1097165.273
	// Create parameters (p, q, d).
	params, err := NewParameters(-3, 1, 16)
	if err != nil {
		t.Error(err)
	}
	// Round up.
	ru := roundUp(r, params)
	if ru != er {
		t.Errorf("expected rational %f but got %f", er, ru)
	}
}
