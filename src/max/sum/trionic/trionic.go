package trionic

const placeholder int64 = -1e14

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maxSumTrionic(nums []int) int64 {
	n := len(nums)
	dp := [4][]int64{}
	for i := 0; i < 4; i++ {
		dp[i] = make([]int64, n)
	}
	dp[0][0] = int64(nums[0])
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[0][i] = max(dp[0][i-1], 0) + int64(nums[i])
		} else {
			dp[0][i] = int64(nums[i])
		}
	}
	for i := 1; i < 4; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = placeholder
		}
	}
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[1][i] = max(dp[1][i-1], dp[0][i-1]) + int64(nums[i])
		}
	}
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			dp[2][i] = max(dp[2][i-1], dp[1][i-1]) + int64(nums[i])
		}
	}
	out := placeholder
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[3][i] = max(dp[3][i-1], dp[2][i-1]) + int64(nums[i])
			if dp[3][i] > out {
				out = dp[3][i]
			}
		}
	}
	return out
}
