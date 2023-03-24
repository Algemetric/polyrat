package sim2d

import (
	"math/big"
	"testing"
)

func TestSymmetricModulo(t *testing.T) {
	// Input values: fraction number (n), radix (r), expected modulo (em).
	n := big.NewRat(-44979, 2401)
	// Expected symmetric modulos.
	var esm []*big.Int
	sm := []int64{0, 3, 3, -46, 640, 5442, -44979}
	for i := 0; i < len(sm); i++ {
		n := big.NewInt(sm[i])
		esm = append(esm, n)
	}
	// Parameters.
	r := []int64{1, 7, 49, 343, 2401, 16807, 117649}
	// Create parameters (b, p, q, d).
	params, err := NewParameters(7, -4, 1, 16)
	if err != nil {
		t.Error(err)
	}
	// Numerator from the given fraction.
	n = isolateNumerator(n, params)
	// Iterate over radix values.
	for i := 0; i < len(r); i++ {
		// Calculates the symmetric modulo.
		m, err := symmetricModulo(n, r[i])
		if err != nil {
			t.Error(err)
		}
		// Check expected symmetric modulo.
		// modulo != expected
		if m.Cmp(esm[i]) == -1 || m.Cmp(esm[i]) == 1 {
			t.Errorf("expected %s but got %s", esm[i].String(), m.String())
		}
	}
}

func TestExpansion(t *testing.T) {
	// Input values: number (n), radix (r).
	f := big.NewRat(-44979, 2401)
	// Create parameters (b, p, q, d).
	params, err := NewParameters(7, -4, 1, 16)
	if err != nil {
		t.Error(err)
	}
	n := isolateNumerator(f, params)
	// Calculate expansion.
	e, err := expansion(n, params)
	if err != nil {
		t.Error(err)
	}
	// Expected expansion.
	ee := []float64{3.0, 0.0, -1.0, 2.0, 2.0, -3.0}
	// Check if calculated expansion matches the expected values of [3, 0, -1, 2, 2, -3].
	if e[0] != ee[0] || e[1] != ee[1] || e[2] != ee[2] || e[3] != ee[3] || e[4] != ee[4] || e[5] != ee[5] {
		t.Errorf("expected expansion of %v but got %v", ee, e)
	}
}
