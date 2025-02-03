package palindrome

func validPalindrome(s string) bool {
	var l int
	r := len(s) - 1
	var option []int
	for l < r {
		if s[l] != s[r] {
			if option == nil {
				option = []int{l, r - 1}
				l = l + 1
				continue
			} else if len(option) != 0 {
				l = option[0]
				r = option[1]
				option = option[:0]
				continue
			} else {
				return false
			}
		}
		l++
		r--
	}
	return true
}
