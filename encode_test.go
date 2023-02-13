package sim2dcodec

import (
	"fmt"
	"math/big"
	"testing"
)

// TestEncode tests the encoding of a fractional number into a polynomial.
func TestEncode(t *testing.T) {
	f := big.NewRat(-44979, 2401)
	b, p, q, d := 7, -4, 1, 8
	// Calculate code.
	c := Encode(f, b, p, q, d)
	// Expected code.
	// [2, -3, 0, 0, -3, 0, 1, -2] = encode(-44979/2401,7,-4,1,8)
	ec := []int64{2, -3, 0, 0, -3, 0, 1, -2}
	// Check results.
	for i := 0; i < len(c); i++ {
		if c[i] != ec[i] {
			t.Errorf("expected %d at position %d but got %d", ec[i], i, c[i])
		}
	}
}

// ExampleEncode runs the case that was initially used to build the encode function.
func ExampleEncode() {
	f := big.NewRat(-44979, 2401)
	b, p, q, d := 7, -4, 1, 8
	c := Encode(f, b, p, q, d)
	fmt.Println(c)
	// Output: [2 -3 0 0 -3 0 1 -2]
}
