package sequence

func validSequence(word1 string, word2 string) []int {
	m := len(word1)
	dp := make([]int, m)
	n := len(word2)
	if word1[m-1] == word2[n-1] {
		dp[m-1] = 1
	}
	for i := m - 2; i >= 0; i-- {
		if dp[i+1] != n && word1[i] == word2[n-1-dp[i+1]] {
			dp[i] = dp[i+1] + 1
		} else {
			dp[i] = dp[i+1]
		}
	}
	out := make([]int, 0, n)
	var changed bool
loop:
	for i := 0; i < m && len(out) < n; i++ {
		if !changed && word1[i] != word2[len(out)] {
			var remain int
			if i+1 < m {
				remain = dp[i+1]
			}
			if remain+len(out)+1 >= n {
				out = append(out, i)
				changed = true
				continue loop
			}
		}
		if word1[i] == word2[len(out)] {
			out = append(out, i)
		}
	}
	if len(out) != n {
		return []int{}
	}
	return out
}
