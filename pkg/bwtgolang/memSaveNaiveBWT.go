package bwtgolang

import (
	"bytes"
	"errors"
	"sort"
)

// CheckEndSymbol is a global variable for checking end symbol before Burrows–Wheeler transform
var CheckEndSymbol = false

// ErrEndSymbolExisted means you should choose another EndSymbol
var ErrEndSymbolExisted = errors.New("bwt: end-symbol existed in string")

func SuffixArray(s []byte) []int {
	sa := make([]int, len(s)+1)
	sa[0] = len(s)

	for i := 0; i < len(s); i++ {
		sa[i+1] = i
	}
	sort.Slice(sa[1:], func(i, j int) bool {
		return bytes.Compare(s[sa[i+1]:], s[sa[j+1]:]) < 0
	})
	return sa
}

var ErrInvalidSuffixArray = errors.New("bwt: invalid suffix array")

func Transform(s []byte) []byte {

	sa := SuffixArray(s)
	bwt := FromSuffixArray(s, sa)
	return bwt
}

// FromSuffixArray compute BWT from sa
func FromSuffixArray(s []byte, sa []int) []byte {
	es := byte('$')
	if len(s)+1 != len(sa) || sa[0] != len(s) {
		return nil
	}
	bwt := make([]byte, len(sa))
	bwt[0] = s[len(s)-1]
	for i := 1; i < len(sa); i++ {
		if sa[i] == 0 {
			bwt[i] = es
		} else {
			bwt[i] = s[sa[i]-1]
		}
	}
	return bwt
}
