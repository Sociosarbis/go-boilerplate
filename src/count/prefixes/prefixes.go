package prefixes

func countPrefixes(words []string, s string) int {
	n1 := len(s)
	var ret int
	for _, word := range words {
		n2 := len(word)
		if n2 <= n1 && word == s[:n2] {
			ret++
		}
	}
	return ret
}
