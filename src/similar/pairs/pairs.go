package pairs

func similarPairs(words []string) int {
	n := len(words)
	counter := make(map[int]int, n)
	for _, word := range words {
		l := len(word)
		var m int
		for i := 0; i < l; i++ {
			m |= 1 << (word[i] - 97)
		}
		if c, ok := counter[m]; ok {
			counter[m] = c + 1
		} else {
			counter[m] = 1
		}
	}
	var ret int
	for _, c := range counter {
		ret += c * (c - 1) / 2
	}
	return ret
}
