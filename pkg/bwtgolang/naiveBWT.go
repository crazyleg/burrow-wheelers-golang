package bwtgolang

import (
	"bytes"
	"sort"
)

func NaiveBWT(data []byte) []byte {
	//start := time.Now()
	var circularArray = make([][]byte, len(data))
	for i := 0; i < len(data); i++ {
		circularArray[i] = make([]byte, len(data))
	}
	//elapsed := time.Since(start)
	//log.Printf("Creation took %s", elapsed)
	//start = time.Now()
	for i := 0; i < len(data); i++ {
		circularArray[i] = append(data[i:], data[:i]...)
	}
	//elapsed = time.Since(start)
	//log.Printf("Rotation took %s", elapsed)

	//start = time.Now()
	sort.Slice(circularArray, func(i, j int) bool {
		return bytes.Compare(circularArray[i], circularArray[j]) < 0
	})
	//elapsed = time.Since(start)
	//log.Printf("sort took %s", elapsed)
	var result []byte = make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = circularArray[i][len(data)-1]
	}

	return result
}
