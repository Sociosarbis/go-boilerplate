package coins

func maxCoins(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 1; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			j := i + l - 1
			var left int
			if i > 0 {
				left = nums[i-1]
			} else {
				left = 1
			}
			var right int
			if j+1 < n {
				right = nums[j+1]
			} else {
				right = 1
			}
			for k := i; k <= j; k++ {
				var leftDp int
				if k-1 <= i {
					leftDp = 1
				} else {
					leftDp = dp[i][k-1]
				}
				var rightDp int
				if k+1 >= j {
					rightDp = 1
				} else {
					rightDp = dp[k+1][j]
				}
				// 当nums[k]是最后一个消时获得的分数，及左右两个区间完全消去获得的分数
				// 巧妙的地方是他会考虑[i,j]区间外紧挨着的左右两个数
				temp := left*right*nums[k] + leftDp + rightDp
				if temp > dp[i][j] {
					dp[i][j] = temp
				}
			}
		}
	}
	return dp[0][n-1]
}
