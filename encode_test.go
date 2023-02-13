package sim2dcodec

import (
	"fmt"
	"math/big"
)

// ExampleEncode runs the case that was initially used to build the encode function.
func ExampleEncode() {
	f := big.NewRat(-44979, 2401)
	b := 7
	p := -4
	q := 1
	d := 8
	c := Encode(f, b, p, q, d)
	fmt.Println(c)
	// Output: [2 -3 0 0 -3 0 1 -2]
}
