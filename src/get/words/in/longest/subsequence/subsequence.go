package subsequence

func diffOne(a, b string) bool {
	var hasDiff bool
	for i := range a {
		if a[i] != b[i] {
			if hasDiff {
				return false
			}
			hasDiff = true
		}
	}
	return hasDiff
}

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	dp := make([][2]int, n)
	var index int
	for i := 0; i < n; i++ {
		dp[i] = [2]int{-1, 1}
		for j := i - 1; j >= 0; j-- {
			if groups[i] != groups[j] && len(words[i]) == len(words[j]) && dp[j][1] >= dp[i][1] && diffOne(words[i], words[j]) {
				dp[i][0], dp[i][1] = j, dp[j][1]+1
				if dp[i][1] > dp[index][1] {
					index = i
				}
			}
		}
	}
	ret := make([]string, dp[index][1])
	i := dp[index][1] - 1
	cur := index
	for cur != -1 {
		ret[i] = words[cur]
		cur = dp[cur][0]
		i--
	}
	return ret
}
