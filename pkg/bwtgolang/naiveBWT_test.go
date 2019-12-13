package bwtgolang

import (
	"testing"
)

func TestNaiveBWT(t *testing.T) {
	r := string(NaiveBWT([]byte("SIX.MIXED.PIXIES.SIFT.SIXTY.PIXIE.DUST.BOXES$")))
	if r != "STEXYDST.E.IXXIIXXSSMPPS.B..EE.$.USFXDIIOIIIT" {
		t.Errorf("NaiveBWT is wrong, %s", r)
	}
}
