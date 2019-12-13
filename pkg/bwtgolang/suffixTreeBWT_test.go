package bwtgolang

import "testing"

func TestSuffixTree_githubBWT(t *testing.T) {
	r := string(SuffixTree_githubBWT([]byte("SIX.MIXED.PIXIES.SIFT.SIXTY.PIXIE.DUST.BOXES$")))
	if string(r) != "STEXYDST.E.IXXIIXXSSMPPS.B..EE.$.USFXDIIOIIIT" {
		t.Errorf("SufTreeBWT  is wrong, %s", r)
	}
}
