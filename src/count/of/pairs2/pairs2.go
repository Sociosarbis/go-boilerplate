package pairs2

const mod int = 1e9 + 7

func countOfPairs(nums []int) int {
	n := len(nums)
	dp := make([][1002]int, n)
	for i := 0; i <= nums[0]; i++ {
		dp[0][i+1] = 1 + dp[0][i]
	}
	for i := nums[0] + 1; i <= 1000; i++ {
		dp[0][i+1] = dp[0][i]
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= nums[i]; j++ {
			k := nums[i] - j
			prevNum := nums[i-1]
			var start int
			if prevNum-k > j {
				start = j
			} else {
				start = prevNum - k
			}
			var delta int
			if start >= 0 {
				delta = dp[i-1][start+1]
			}
			dp[i][j+1] = (dp[i][j] + delta) % mod
		}
		for j := nums[i] + 1; j <= 1000; j++ {
			dp[i][j+1] = dp[i][j]
		}
	}
	return dp[n-1][1001]
}
