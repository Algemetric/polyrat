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
	p := -4
	b := 7
	// Numerator from the given fraction.
	n = isolateNumerator(n, b, p)
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
	r, p, q := 7, -4, 1
	n := isolateNumerator(f, r, p)
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

func TestValidateModulo(t *testing.T) {
	// b >= 2.
	b := 1
	err := validateModulo(b)
	if err == nil {
		t.Error("an error should be thrown when b is less than 2")
	} else {
		if err.Error() != ErrBIsLessThan2.Error() {
			t.Error(ErrBIsLessThan2.Error())
		}
	}
}

func TestValidateSmallestPowerOfExpansion(t *testing.T) {
	// p < q.
	p, q := 1, 1
	err := validateSmallestPowerOfExpansion(p, q)
	if err == nil {
		t.Error("an error should be thrown when p is less than q")
	} else {
		if err.Error() != ErrPIsLessThanQ.Error() {
			t.Error(ErrPIsLessThanQ.Error())
		}
	}
	// p < 0.
	p, q = 0, 2
	err = validateSmallestPowerOfExpansion(p, q)
	if err == nil {
		t.Error("an error should be thrown when p is greater than or equal to 0")
	} else {
		if err.Error() != ErrPIsGreaterThanOrEqualToZero.Error() {
			t.Error(ErrPIsGreaterThanOrEqualToZero.Error())
		}
	}
}

func TestValidateGreatestPowerOfExpansion(t *testing.T) {
	// q > 0.
	q := 0
	err := validateGreatestPowerOfExpansion(q)
	if err == nil {
		t.Error("an error should be thrown when q is less than or equal to 0")
	} else {
		if err.Error() != ErrQIsLessThanOrEqualToZero.Error() {
			t.Error(ErrQIsLessThanOrEqualToZero.Error())
		}
	}
}

func TestValidateDegree(t *testing.T) {
	// d is a power of 2.
	p, q, d := -4, 1, 3
	err := validateDegree(p, q, d)
	if err == nil {
		t.Error("an error should be thrown when d is not a power of 2")
	} else {
		if err.Error() != ErrDIsNotAPowerOf2.Error() {
			t.Error(ErrDIsNotAPowerOf2.Error())
		}
	}
	// d >= 1.
	p, q, d = -4, 1, 0
	err = validateDegree(p, q, d)
	if err == nil {
		t.Error("an error should be thrown when d is less than 1")
	} else {
		if err.Error() != ErrDIsLessThanOne.Error() {
			t.Error(ErrDIsLessThanOne.Error())
		}
	}
	// d > q + |p|.
	p, q, d = 1, 7, 8
	err = validateDegree(p, q, d)
	if err == nil {
		t.Error("an error should be thrown when d is less than or equal to q plus the absolute value of p")
	} else {
		if err.Error() != ErrDIsLessThanOrEqualToQPlusP.Error() {
			t.Error(ErrDIsLessThanOrEqualToQPlusP.Error())
		}
	}
}

func TestValidateParameters(t *testing.T) {
	// Valid parameters.
	b, p, q, d := 7, -4, 1, 8
	err := validateParameters(b, p, q, d)
	if err != nil {
		t.Error("parameters should be valid")
	}
}
