package minimum

func addMinimum(word string) int {
	var s byte = 0
	ret := 0
	i := 0
	for i < len(word) {
		if word[i]-97 == s {
			i++
		} else {
			ret++
		}
		s = (s + 1) % 3
	}
	if s == 0 {
		return ret
	}
	return ret + int(3-s)
}
