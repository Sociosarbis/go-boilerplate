package checker

type node struct {
	has_end bool
	next    []*node
}

type StreamChecker struct {
	trie   *node
	suffix []byte
}

func Constructor(words []string) StreamChecker {
	trie := node{
		false,
		make([]*node, 26),
	}
	maxLen := 0
	for _, word := range words {
		n := &trie
		if len(word) > maxLen {
			maxLen = len(word)
		}
		for i := len(word) - 1; i >= 0; i-- {
			index := word[i] - 97
			if n.next[index] == nil {
				n.next[index] = &node{
					false,
					make([]*node, 26),
				}
			}
			n = n.next[index]
		}
		n.has_end = true
	}
	return StreamChecker{
		&trie,
		make([]byte, maxLen),
	}
}

func (this *StreamChecker) Query(letter byte) bool {
	copy(this.suffix[:], this.suffix[1:])
	this.suffix[len(this.suffix)-1] = letter
	node := this.trie
	for i := len(this.suffix) - 1; i >= 0; i-- {
		if this.suffix[i] == 0 {
			break
		}
		index := this.suffix[i] - 97
		if node.next[index] == nil {
			break
		}
		if node.next[index].has_end {
			return true
		}
		node = node.next[index]
	}
	return false
}
