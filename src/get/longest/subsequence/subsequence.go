package subsequence

func getLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	ret := make([]string, 0, n)
	ret = append(ret, words[0])
	for i := 1; i < n; i++ {
		if groups[i] != groups[i-1] {
			ret = append(ret, words[i])
		}
	}
	return ret
}
