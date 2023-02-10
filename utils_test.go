package sim2dcodec

import (
	"math/big"
	"testing"
)

func TestSymmetricModulo(t *testing.T) {
	// Input values: number (n), radix (r), expected modulo (em).
	n := int64(-44979)
	r := int64(1)
	// Calculates the symmetric modulo.
	m, err := symmetricModulo(n, r)
	if err != nil {
		t.Error(err)
	}
	// Check expected symmetric modulo.
	em := 0
	if m != 0 {
		t.Errorf("expected %d but got %d", em, m)
	}
	// Value for radix should trigger an error.
	// n = -44979
	// r = 0
	// Calculates the symmetric modulo.
	// Checks if function throwns an error.
	// m = symmetricModulo(n, r)
}

func TestExpansion(t *testing.T) {
	// Input values: number (n), radix (r).
	n := int64(-44979)
	r := 7
	p := -4
	q := 1
	pl := polynomialLength(q, p)
	// Calculate expansion.
	e := expansion(pl, r, n)
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

func TestEncode(t *testing.T) {
	// Expected result.
	// [2, -3, 0, 0, -3, 0, 1, -2] = encode(-44979/2401,7,-4,1,8)
	fraction := big.NewRat(-44979, 2401)
	base := 7
	p := -4
	q := 1
	d := 8
	// Calculate code.
	c := Encode(fraction, base, p, q, d)
	// Expected code.
	ec := []int64{2, -3, 0, 0, -3, 0, 1, -2}
	// Check results.
	if ec[0] != c[0] || ec[1] != c[1] || ec[2] != c[2] || ec[3] != c[3] || ec[4] != c[4] || ec[5] != c[5] || ec[6] != c[6] || ec[7] != c[7] {
		t.Errorf("expected %v but got %v", ec, c)
	}
}
