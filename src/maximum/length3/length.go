package length3

func maximumLength(nums []int) int {
	dp := [2][2]int{}
	n := len(nums)
	i := nums[0] & 1
	dp[i][0], dp[i][1] = 1, 1
	for i := 1; i < n; i++ {
		i := nums[i] & 1
		dp[i][0]++
		if i == 1 || dp[1-i][1] != 0 {
			dp[i][1] = dp[1-i][1] + 1
		}
	}
	var ret int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if dp[i][j] > ret {
				ret = dp[i][j]
			}
		}
	}
	return ret
}
