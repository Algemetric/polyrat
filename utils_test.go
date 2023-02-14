package sim2dcodec

import (
	"math/big"
	"testing"
)

func TestSymmetricModulo(t *testing.T) {
	// Input values: number (n), radix (r), expected modulo (em).
	n := big.NewRat(-44979, 1)
	r := int64(1)
	// Calculates the symmetric modulo.
	m, err := symmetricModulo(n, r)
	if err != nil {
		t.Error(err)
	}
	// Check expected symmetric modulo.
	em := 0
	if m.Cmp(big.NewInt(0)) == -1 || m.Cmp(big.NewInt(0)) == 1 {
		t.Errorf("expected %d but got %d", em, m)
	}
	// Value for radix = 0 should trigger an error.
	r = int64(0)
	// Calculates the symmetric modulo.
	// Checks if function throws an error.
	m, err = symmetricModulo(n, r)
	if err == nil {
		t.Error("radix 0 should throw an error")
	}
}

func TestExpansion(t *testing.T) {
	t.Skip()
	// Input values: number (n), radix (r).
	n := big.NewRat(-44979, 2401)
	r, p, q := 7, -4, 1
	pl := polynomialLength(q, p)
	// Calculate expansion.
	e, err := expansion(pl, r, n)
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

func TestPolynomialLength(t *testing.T) {
	t.Skip()
}
