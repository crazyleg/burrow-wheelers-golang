package bwtgolang

import (
	"testing"
)

func TestSuffixArrayBWT(t *testing.T) {
	r := string(SuffixArrayBWT([]byte("SIX.MIXED.PIXIES.SIFT.SIXTY.PIXIE.DUST.BOXES$")))
	if string(r) != "STEXYDST.E.IXXIIXXSSMPPS.B..EE.$.USFXDIIOIIIT" {

		t.Errorf("SufTreeBWT  is wrong, %s", r)
	}
}
