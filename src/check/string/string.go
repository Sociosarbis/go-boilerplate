package string

func checkString(s string) bool {
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i] == 'a' && s[i-1] == 'b' {
			return false
		}
	}
	return true
}
