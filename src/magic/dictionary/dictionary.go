package dictionary

type MagicDictionary struct {
	hasEnd bool
	next   []*MagicDictionary
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		hasEnd: false,
		next:   make([]*MagicDictionary, 26),
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		cur := this
		for i := 0; i < len(word); i++ {
			index := word[i] - 97
			if cur.next[index] == nil {
				item := Constructor()
				cur.next[index] = &item
			}
			cur = cur.next[index]
		}
		cur.hasEnd = true
	}
}

func dfs(node *MagicDictionary, searchWord string, i int, hasReplaced bool) bool {
	index := int(searchWord[i] - 97)
	if i+1 == len(searchWord) {
		if node.next[index] != nil {
			if node.next[index].hasEnd && hasReplaced {
				return true
			}
		}
		if !hasReplaced {
			for idx, temp := range node.next {
				if temp != nil && idx != index {
					if temp.hasEnd {
						return true
					}
				}
			}
		}
		return false
	}
	if node.next[index] == nil {
		if hasReplaced {
			return false
		}
	} else {
		temp := node.next[index]
		ret := dfs(temp, searchWord, i+1, hasReplaced)
		if ret {
			return ret
		}
	}

	if !hasReplaced {
		for idx, temp := range node.next {
			if temp != nil && idx != index {
				if dfs(temp, searchWord, i+1, true) {
					return true
				}
			}
		}
	}
	return false
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return dfs(this, searchWord, 0, false)
}
