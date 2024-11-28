package pairs

const mod int = 1e9 + 7

func countOfPairs(nums []int) int {
	n := len(nums)
	dp := make([][51][51]int, n)
	for i := 0; i <= nums[0]; i++ {
		dp[0][i][nums[0]-i] = 1
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
			for jj := start; jj >= 0; jj-- {
				kk := prevNum - jj
				if dp[i-1][jj][kk] != 0 {
					dp[i][j][k] = (dp[i][j][k] + dp[i-1][jj][kk]) % mod
				}
			}
		}
	}
	var ret int
	for _, pairs := range dp[n-1] {
		for _, v := range pairs {
			if v != 0 {
				ret = (ret + v) % mod
			}
		}
	}
	return ret
}
