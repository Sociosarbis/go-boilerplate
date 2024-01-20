package separator

func splitWordsBySeparator(words []string, separator byte) []string {
	ret := make([]string, 0, len(words))
	for _, word := range words {
		start := -1
		n := len(word)
		for j := 0; j < n; j++ {
			if word[j] == separator {
				if j-start > 1 {
					ret = append(ret, word[start+1:j])
				}
				start = j
			}
		}
		if start+1 != n {
			ret = append(ret, word[start+1:])
		}
	}
	return ret
}
