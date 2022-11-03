package repeating

func maxRepeating(sequence string, word string) int {
	if len(sequence) < len(word) {
		return 0
	}
	dp := make([]int, len(sequence)-len(word)+1)
	ret := 0
	for i := 0; i < len(dp); i++ {
		if sequence[i:i+len(word)] == word {
			if i >= len(word) {
				dp[i] = dp[i-len(word)] + 1
			} else {
				dp[i] = 1
			}
			if dp[i] > ret {
				ret = dp[i]
			}
		}
	}
	return ret
}
