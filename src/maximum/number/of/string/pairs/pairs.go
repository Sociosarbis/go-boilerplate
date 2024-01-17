package pairs

func maximumNumberOfStringPairs(words []string) int {
	dict := make(map[string]struct{}, len(words))
	ret := 0
	for _, word := range words {
		if _, ok := dict[word]; ok {
			ret++
		} else {
			dict[reverseWord(word)] = struct{}{}
		}
	}
	return ret
}

func reverseWord(word string) string {
	ret := []byte(word)
	l := 0
	r := len(ret) - 1
	for l < r {
		ret[l], ret[r] = ret[r], ret[l]
		l++
		r--
	}
	return string(ret)
}
