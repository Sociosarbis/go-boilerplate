package length4

func maximumLength(nums []int, k int) int {
	dp := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, k)
	}
	n := len(nums)
	i := nums[0] % k
	for j := 0; j < k; j++ {
		dp[i][j] = 1
	}
	for i := 1; i < n; i++ {
		i := nums[i] % k
		for j := 0; j < k; j++ {
			prev := (j + k - i) % k
			if dp[i][j] == 0 || dp[prev][j] != 0 {
				dp[i][j] = dp[prev][j] + 1
			}
		}
	}
	var ret int
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			if dp[i][j] > ret {
				ret = dp[i][j]
			}
		}
	}
	return ret
}
