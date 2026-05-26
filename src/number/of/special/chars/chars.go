package chars

func numberOfSpecialChars(word string) int {
	set := [26]bool{}
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			set[c-'A'] = true
		}
	}
	var out int
	for _, c := range word {
		if c >= 'a' && c <= 'z' {
			if set[c-'a'] {
				out++
				set[c-'a'] = false
			}
		}
	}
	return out
}
