package bwtgolang

import (
	"testing"
)

func TestTransform(t *testing.T) {
	r := string(Transform([]byte("^BANANA")))
	if r != "ANNB^AA$" {
		t.Errorf("Mem Efficient  is wrong, %s", r)
	}

}
