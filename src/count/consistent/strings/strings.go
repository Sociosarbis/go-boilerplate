package strings

func countConsistentStrings(allowed string, words []string) int {
	m := make([]bool, 26)
	for _, c := range allowed {
		m[c-97] = true
	}
	ret := 0
loop:
	for _, word := range words {
		for _, c := range word {
			if !m[c-97] {
				continue loop
			}
		}
		ret++
	}
	return ret
}
