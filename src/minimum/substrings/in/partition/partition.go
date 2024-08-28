package partition

func minimumSubstringsInPartition(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
	}
	for i := 0; i < n; i++ {
		counter := [26]int{}
		uniqCounter := map[int]int{}
		for j := i; j < n; j++ {
			index := s[j] - 97
			oldValue := counter[index]
			if oldValue != 0 {
				if c, ok := uniqCounter[oldValue]; ok {
					if c == 1 {
						delete(uniqCounter, oldValue)
					} else {
						uniqCounter[oldValue] = c - 1
					}
				}
			}
			if c, ok := uniqCounter[oldValue+1]; ok {
				uniqCounter[oldValue+1] = c + 1
			} else {
				uniqCounter[oldValue+1] = 1
			}
			counter[index] = oldValue + 1
			if len(uniqCounter) == 1 {
				if dp[i]+1 < dp[j+1] {
					dp[j+1] = dp[i] + 1
				}
			}
		}
	}
	return dp[n]
}
