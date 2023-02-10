package sim2dcodec

import (
	"math/big"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	f := big.NewRat(-44979, 2401)
	b := 7
	p := -4
	q := 1
	d := 8
	c := Encode(f, b, p, q, d)
	rf := Decode(c, b, p, q)
	if strings.Compare(f.String(), rf.String()) != 0 {
		t.Errorf("error decoding")
	}
}
