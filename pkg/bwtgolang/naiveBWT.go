package bwtgolang

import (
	"bytes"
	"sort"
)

func NaiveBWT(data []byte) []byte {

	var circularArray = make([][]byte, len(data))
	for i := 0; i < len(data); i++ {
		circularArray[i] = make([]byte, len(data))
	}
	for i := 0; i < len(data); i++ {
		circularArray[i] = append(data[i:], data[:i]...)
	}

	sort.Slice(circularArray, func(i, j int) bool {
		return bytes.Compare(circularArray[i], circularArray[j]) < 0
	})

	var result []byte = make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = circularArray[i][len(data)-1]
	}

	return result
}
