package char

type trie struct {
	exists bool
	next   []*trie
}

func minExtraChar(s string, dictionary []string) int {
	root := &trie{}
	n := len(s)
	for _, word := range dictionary {
		cur := root
		for i := 0; i < len(word); i++ {
			if cur.next == nil {
				cur.next = make([]*trie, 26)
			}
			if cur.next[word[i]-97] == nil {
				cur.next[word[i]-97] = &trie{}
			}
			cur = cur.next[word[i]-97]
		}
		cur.exists = true
	}
	ret := make([]int, n)
	for i := range ret {
		ret[i] = i + 1
	}
	for i := 0; i < n; i++ {
		cur := root
		var prev int
		if i > 0 {
			prev = ret[i-1]
		}
		if prev+1 < ret[i] {
			ret[i] = prev + 1
		}
		for j := 0; i+j < n; j++ {
			if cur.next == nil || cur.next[s[i+j]-97] == nil {
				break
			}
			cur = cur.next[s[i+j]-97]
			if cur.exists {
				if ret[i+j] > prev {
					ret[i+j] = prev
				}
			}
		}
	}
	return ret[n-1]
}
