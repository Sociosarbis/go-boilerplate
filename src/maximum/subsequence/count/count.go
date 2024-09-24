package count

func maximumSubsequenceCount(text string, pattern string) int64 {
	a, b := pattern[0], pattern[1]
	var count int64
	var t1 int64
	n := len(text)
	for i := n - 1; i >= 0; i-- {
		if text[i] == a {
			t1 += count
		}
		if text[i] == b {
			count++
		}
	}
	t1 += count
	count = 0
	var t2 int64
	for i := 0; i < n; i++ {
		if text[i] == b {
			t2 += count
		}
		if text[i] == a {
			count++
		}
	}
	t2 += count
	if t1 > t2 {
		return t1
	}
	return t2
}
