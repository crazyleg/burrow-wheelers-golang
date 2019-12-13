package bwtgolang

import (
	"index/suffixarray"
	"reflect"
)

func SuffixArrayBWT(s []byte) []byte {
	sa := suffixarray.New(s)
	bwt := make([]byte, len(s))

	tmp := reflect.ValueOf(sa).Elem().FieldByName("sa").FieldByIndex([]int{0})
	var sa1 []int = make([]int, len(s))

	for i := 0; i < len(s); i++ {
		sa1[i] = int(tmp.Index(i).Int())
	}

	for i := 0; i < len(s); i++ {
		if sa1[i] == 0 {
			bwt[i] = '$'
		} else {
			bwt[i] = s[sa1[i]-1]
		}

	}
	//fmt.Println(string(bwt))
	return bwt
}
