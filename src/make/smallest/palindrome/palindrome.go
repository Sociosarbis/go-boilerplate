package palindrome

func makeSmallestPalindrome(s string) string {
	l := 0
	r := len(s) - 1
	bytes := []byte(s)
	for l < r {
		if bytes[l] < bytes[r] {
			bytes[r] = bytes[l]
		} else if bytes[l] > bytes[r] {
			bytes[l] = bytes[r]
		}
		l++
		r--
	}
	return string(bytes)
}
