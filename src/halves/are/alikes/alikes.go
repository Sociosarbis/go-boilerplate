package alikes

func isVowel(c byte) bool {
	if c >= 65 && c <= 90 {
		c += 32
	}
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func halvesAreAlike(s string) bool {
	n := len(s) / 2
	i := 0
	count := 0
	for ; i < n; i++ {
		if isVowel(s[i]) {
			count++
		}
	}
	for ; i < len(s); i++ {
		if isVowel(s[i]) {
			count--
		}
	}
	return count == 0
}
