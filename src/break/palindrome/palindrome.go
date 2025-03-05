package palindrome

func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	end := n / 2
	s := []byte(palindrome)
	for i := 0; i < end; i++ {
		if palindrome[i] != 'a' {
			s[i] = 'a'
			return string(s)
		}
	}
	if n > 1 {
		s[n-1]++
		return string(s)
	}
	return ""
}
