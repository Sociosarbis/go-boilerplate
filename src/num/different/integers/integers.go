package integers

func numDifferentIntegers(word string) int {
	m := make(map[string]struct{}, 0)
	l := 0
	for i, c := range word {
		if c >= 97 && c <= 122 {
			if l < i {
				for l+1 < i && word[l] == '0' {
					l++
				}
				if _, ok := m[word[l:i]]; !ok {
					m[word[l:i]] = struct{}{}
				}
			}
			l = i + 1
		}
	}
	ret := len(m)
	if l < len(word) {
		for l+1 < len(word) && word[l] == '0' {
			l++
		}
		if _, ok := m[word[l:]]; !ok {
			ret++
		}
	}

	return ret
}
