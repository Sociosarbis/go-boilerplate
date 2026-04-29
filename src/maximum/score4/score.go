package score4

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maximumScore(grid [][]int) int64 {
	n := len(grid)
	if n == 1 {
		return 0
	}
	// 先算每一列的前缀和
	colSum := make([][]int64, n)
	for i := 0; i < n; i++ {
		colSum[i] = make([]int64, n+1)
		for j := 0; j < n; j++ {
			colSum[i][j+1] = colSum[i][j] + int64(grid[j][i])
		}
	}
	dp := make([][][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int64, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int64, n+1)
		}
	}
	// 在上一列高度为prevH的条件下，上上列的高度在0到j时，去除上一列的得分后（为了不重复计算），所能获得的最大值
	prevMax := make([][]int64, n+1)
	// 在上一列高度为prevH的条件下，上上列的高度在j到n时所能获得的最大值
	prevSuffixMax := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		prevMax[i] = make([]int64, n+1)
		prevSuffixMax[i] = make([]int64, n+1)
	}

	for i := 1; i < n; i++ {
		for currH := 0; currH <= n; currH++ {
			for prevH := 0; prevH <= n; prevH++ {
				if currH <= prevH {
					extraScore := colSum[i][prevH] - colSum[i][currH]
					dp[i][currH][prevH] = prevSuffixMax[prevH][0] + extraScore
				} else {
					extraScore := colSum[i-1][currH] - colSum[i-1][prevH]
					// 上上列的高度不小于当前列时
					// 或上上列的高度不大于当前列，所以可以加上extraScore
					dp[i][currH][prevH] = max(prevSuffixMax[prevH][currH], prevMax[prevH][currH]+extraScore)
				}
			}
		}
		for currH := 0; currH <= n; currH++ {
			// 因为不用扣除当前列得分
			prevMax[currH][0] = dp[i][currH][0]
			for prevH := 1; prevH <= n; prevH++ {
				var penalty int64
				// 去掉当前列得分
				if prevH > currH {
					penalty = colSum[i][prevH] - colSum[i][currH]
				}
				prevMax[currH][prevH] = max(prevMax[currH][prevH-1], dp[i][currH][prevH]-penalty)
			}
			prevSuffixMax[currH][n] = dp[i][currH][n]
			for prevH := n - 1; prevH >= 0; prevH-- {
				prevSuffixMax[currH][prevH] = max(prevSuffixMax[currH][prevH+1], dp[i][currH][prevH])
			}
		}
	}
	// 最后一列，只可能在高度0和满高的情况下得到最大值
	var ans int64
	for k := 0; k <= n; k++ {
		ans = max(ans, max(dp[n-1][0][k], dp[n-1][n][k]))
	}
	return ans
}
