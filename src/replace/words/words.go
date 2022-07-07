package words

type trie struct {
	hasEnd bool
	next   map[byte]*trie
}

func replaceWords(dictionary []string, sentence string) string {
	root := &trie{
		false,
		make(map[byte]*trie, 0),
	}
loop:
	for _, word := range dictionary {
		cur := root
		for i := 0; i < len(word); i++ {
			if cur.hasEnd {
				continue loop
			} else {
				if _, ok := cur.next[word[i]]; !ok {
					cur.next[word[i]] = &trie{
						false,
						make(map[byte]*trie, 0),
					}
				}
				cur = cur.next[word[i]]
			}
		}
		cur.hasEnd = true
		cur.next = make(map[byte]*trie, 0)
	}

	ret := ""
	for i := 0; i < len(sentence); i++ {
		if sentence[i] != ' ' {
			j := i
			for j < len(sentence) && sentence[j] != ' ' {
				j++
			}
			cur := root
			k := i
			for i < j {
				if nextCur, ok := cur.next[sentence[i]]; ok {
					cur = nextCur
					if cur.hasEnd {
						ret += sentence[k : i+1]
						break
					}
				} else {
					ret += sentence[k:j]
					break
				}
				i++
			}
			if i == j {
				ret += sentence[k:j]
			}
			i = j - 1
		} else {
			ret += " "
		}
	}
	return ret
}
