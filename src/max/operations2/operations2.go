package operations2

func newMatrix(n int) [][]int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	return dp
}

func maxOperations(nums []int) int {
	n := len(nums)
	if n == 2 {
		return 1
	}
	ret := 1
	options := [3][4]int{{0, 1, 2, n - 1}, {n - 2, n - 1, 0, n - 3}, {0, n - 1, 1, n - 2}}
	for _, option := range options {
		dp := newMatrix(n)
		target := nums[option[0]] + nums[option[1]]
		dp[option[2]][option[3]] = 1
		for i := option[2]; i+1 < n; i++ {
			for j := option[3]; j >= i+1; j-- {
				if dp[i][j] != 0 {
					if nums[i]+nums[i+1] == target {
						if dp[i][j]+1 > ret {
							ret = dp[i][j] + 1
						}
						if i+2 <= j {
							dp[i+2][j] = dp[i][j] + 1
						}
					}
					if nums[j-1]+nums[j] == target {
						if dp[i][j]+1 > ret {
							ret = dp[i][j] + 1
						}
						if i <= j-2 {
							dp[i][j-2] = dp[i][j] + 1
						}
					}
					if nums[i]+nums[j] == target {
						if dp[i][j]+1 > ret {
							ret = dp[i][j] + 1
						}
						if i+1 <= j-1 {
							dp[i+1][j-1] = dp[i][j] + 1
						}
					}
				}
			}
		}
	}
	return ret
}
