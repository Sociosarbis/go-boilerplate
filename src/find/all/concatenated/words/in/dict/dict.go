package dict

import "sort"

type Trie struct {
	hasEnd   bool
	branches []*Trie
}

type Words []string

func (this Words) Len() int {
	return len(this)
}

func (this Words) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Words) Less(i, j int) bool {
	return len(this[i]) < len(this[j])
}

func findAllConcatenatedWordsInADict(words []string) []string {
	root := Trie{
		hasEnd:   false,
		branches: make([]*Trie, 26),
	}
	sort.Sort(Words(words))
	ret := []string{}
	for _, word := range words {
		res := dfs(word, 0, &root, &root, true)
		if res {
			ret = append(ret, word)
		}
	}
	return ret
}

func dfs(word string, i int, trie *Trie, root *Trie, isInsert bool) bool {
	if i < len(word) {
		b := word[i] - 97
		if trie.branches[b] == nil {
			if isInsert {
				trie.branches[b] = &Trie{
					hasEnd:   i+1 == len(word),
					branches: make([]*Trie, 26),
				}
				dfs(word, i+1, trie.branches[b], root, isInsert)
			}
			return false
		} else {
			if trie.branches[b].hasEnd {
				if i+1 == len(word) {
					return true
				}
				res := dfs(word, i+1, root, root, false)
				if res {
					return true
				}
			}
			res := dfs(word, i+1, trie.branches[b], root, isInsert)
			if res {
				return true
			}
		}
	}
	return false
}
