package filter

type trie struct {
	index int
	next  []*trie
}

type WordFilter struct {
	words []string
	trie  *trie
}

func Constructor(words []string) WordFilter {
	root := &trie{
		index: -1,
		next:  make([]*trie, 26),
	}

	for index, word := range words {
		cur := root
		for i := 0; i < len(word); i++ {
			if cur.next[word[i]-97] == nil {
				cur.next[word[i]-97] = &trie{
					index: -1,
					next:  make([]*trie, 26),
				}
			}
			cur = cur.next[word[i]-97]
		}
		cur.index = index
	}
	return WordFilter{
		words: words,
		trie:  root,
	}
}

func (this *WordFilter) dfs(node *trie, pref []byte, suff string) int {
	ret := -1
	for i := 0; i < len(node.next); i++ {
		if node.next[i] != nil {
			pref = append(pref, byte(97+i))
			temp := this.dfs(node.next[i], pref, suff)
			pref = pref[:len(pref)-1]
			if temp != -1 {
				if temp > ret {
					ret = temp
				}
			}
		}
	}
	if node.index != -1 {
		if node.index > ret {
			if string(pref)[len(pref)-len(suff):] == suff {
				ret = node.index
			}
		}
	}
	return ret
}

func (this *WordFilter) F(pref string, suff string) int {
	cur := this.trie
	for i := 0; i < len(pref); i++ {
		index := pref[i] - 97
		if cur.next[index] != nil {
			cur = cur.next[index]
		} else {
			return -1
		}
	}
	return this.dfs(cur, []byte(pref), suff)
}
