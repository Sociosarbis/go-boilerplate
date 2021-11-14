package sum

type Trie struct {
	Val      int
	Branches []*Trie
}

type MapSum struct {
	Root *Trie
}

func Constructor() MapSum {
	return MapSum{
		Root: &Trie{
			0,
			make([]*Trie, 26),
		},
	}
}

func (this *MapSum) Insert(key string, val int) {
	node := this.Root
	for _, c := range key {
		i := c - 97
		if node.Branches[i] == nil {
			node.Branches[i] = &Trie{
				0,
				make([]*Trie, 26),
			}
		}
		node = node.Branches[i]
	}
	node.Val = val
}

func (this *MapSum) Sum(prefix string) int {
	node := this.Root
	for _, c := range prefix {
		i := c - 97
		if node.Branches[i] == nil {
			return 0
		}
		node = node.Branches[i]
	}
	return Dfs(node)
}

func Dfs(trie *Trie) int {
	ret := trie.Val
	for _, t := range trie.Branches {
		if t != nil {
			ret += Dfs(t)
		}
	}
	return ret
}
