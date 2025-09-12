package win

func isVowel(b byte) bool {
	if b < 97 {
		b += 32
	}
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func doesAliceWin(s string) bool {
	n := len(s)
	for i := 0; i < n; i++ {
		if isVowel(s[i]) {
			return true
		}
	}
	return false
}
