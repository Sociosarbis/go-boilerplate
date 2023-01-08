package count

func prefixCount(words []string, pref string) int {
	n := len(pref)
	ret := 0
	for _, word := range words {
		if len(word) >= n && word[:n] == pref {
			ret++
		}
	}
	return ret
}
