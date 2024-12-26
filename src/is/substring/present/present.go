package present

func isSubstringPresent(s string) bool {
	n := len(s)
	m := make([]bool, 1024)
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			return true
		}
		m[((int(s[i])-97)<<5)|(int(s[i-1])-97)] = true
	}
	for i := n - 2; i >= 0; i-- {
		if m[(int(s[i]-97)<<5)|(int(s[i+1])-97)] {
			return true
		}
	}
	return false
}
