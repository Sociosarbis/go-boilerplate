package strings

func numOfStrings(patterns []string, word string) int {
	var out int
	n := len(word)
	for _, pattern := range patterns {
		m := len(pattern)
		if m <= n {
			for i := n - m; i >= 0; i-- {
				if pattern == word[i:i+m] {
					out++
					break
				}
			}
		}
	}
	return out
}
