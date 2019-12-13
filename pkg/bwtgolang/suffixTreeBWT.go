package bwtgolang

import (
	"sort"
)

type tree []node

type node struct {
	sub string // a substring of the input string
	ch  []int  // list of child nodes
}

func buildTree(s string) tree {
	t := tree{node{}} // root node
	for i := range s {
		t = t.addSuffix(s[i:])
	}
	return t
}

func (t tree) addSuffix(suf string) tree {
	n := 0
	for i := 0; i < len(suf); {
		b := suf[i]
		ch := t[n].ch
		var x2, n2 int
		for ; ; x2++ {
			if x2 == len(ch) {
				// no matching child, remainder of suf becomes new node.
				n2 = len(t)
				t = append(t, node{sub: suf[i:]})
				t[n].ch = append(t[n].ch, n2)
				return t
			}
			n2 = ch[x2]
			if t[n2].sub[0] == b {
				break
			}
		}
		// find prefix of remaining suffix in common with child
		sub2 := t[n2].sub
		j := 0
		for ; j < len(sub2); j++ {
			if suf[i+j] != sub2[j] {
				// split n2
				n3 := n2
				// new node for the part in common
				n2 = len(t)
				t = append(t, node{sub2[:j], []int{n3}})
				t[n3].sub = sub2[j:] // old node loses the part in common
				t[n].ch[x2] = n2
				break // continue down the tree
			}
		}
		i += j // advance past part in common
		n = n2 // continue down the tree
	}
	return t
}

var counter = 0

func sortTree(t tree, n int, curlen int, curstring string, result *[]byte, s *[]byte) {
	if len(t[n].ch) > 1 {
		sort.Slice(t[n].ch, func(i, j int) bool {
			return t[t[n].ch[i]].sub < t[t[n].ch[j]].sub
		})
	}
	for i := range t[n].ch {
		sortTree(t, t[n].ch[i], curlen+len(t[n].sub), curstring+t[n].sub, result, s)
	}
	if len(t[n].sub) > 0 && t[n].sub[len(t[n].sub)-1] == '$' {
		if len(*s)-(curlen+len(t[n].sub))-1 >= 0 {
			(*result)[counter] = (*s)[len(*s)-(curlen+len(t[n].sub))-1]
		} else {
			(*result)[counter] = '$'
		}
		counter++
	}

}

func SuffixTreeBWT(s []byte) []byte {
	var result []byte = make([]byte, len(s))
	counter = 0
	tree := buildTree(string(s))
	sortTree(tree, 0, 0, "", &result, &s)
	//parseTree(tree, 0, 0, "", &result, &s)

	return result
}
