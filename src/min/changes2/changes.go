package changes2

func minChanges(s string) int {
	n := len(s)
	var out int
	for i := 0; i < n; i += 2 {
		if s[i] != s[i+1] {
			out++
		}
	}
	return out
}
