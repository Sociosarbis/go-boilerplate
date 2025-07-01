package count

func possibleStringCount(word string) int {
	n := len(word)
	count := 1
	var out int
	for i := 1; i < n; i++ {
		if word[i] != word[i-1] {
			out += count - 1
			count = 1
		} else {
			count++
		}
	}
	return out + count
}
