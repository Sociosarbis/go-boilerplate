package jumps

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, n)
	var t int64 = int64(target)
	for i := 1; i < n; i++ {
		num := int64(nums[i])
		dp[i] = -1
		for j := 0; j < i; j++ {
			if dp[j] != -1 && num <= int64(nums[j])+t && num >= int64(nums[j])-t {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}
