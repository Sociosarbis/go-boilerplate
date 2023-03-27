package substrings

type node struct {
	count int
	next  []*node
}

func countSubstrings(s string, t string) int {
	trie := &node{
		0,
		make([]*node, 26),
	}
	for i := range t {
		n := trie
		for j := i; j < len(t); j++ {
			index := t[j] - 97
			if n.next[index] == nil {
				n.next[index] = &node{
					0,
					make([]*node, 26),
				}
			}
			n.next[index].count++
			n = n.next[index]
		}
	}

	ret := 0
	for i := range s {
		for j := i; j < len(s); j++ {
			ret += dfs(s, i, j, trie, false)
		}
	}
	return ret
}

func dfs(s string, i, j int, trie *node, changed bool) int {
	if i > j {
		if changed && trie != nil {
			return trie.count
		} else {
			return 0
		}
	}

	ret := 0

	index := byte(s[i] - 97)

	if trie.next[index] != nil {
		ret += dfs(s, i+1, j, trie.next[index], changed)
	}

	if !changed {
		for k := byte(0); k < 26; k++ {
			if trie.next[k] != nil && k != index {
				ret += dfs(s, i+1, j, trie.next[k], true)
			}
		}
	}

	return ret
}
