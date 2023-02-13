package sim2dcodec

import (
	"math/big"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	f := big.NewRat(-44979, 2401)
	b, p, q, d := 7, -4, 1, 8
	c, err := Encode(f, b, p, q, d)
	if err != nil {
		t.Error(err)
	}
	rf := Decode(c, b, p, q)
	if strings.Compare(f.String(), rf.String()) != 0 {
		t.Errorf("error decoding, expected %s but got %s", f.String(), rf.String())
	}
}
