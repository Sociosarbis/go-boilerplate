package walls

var max int = 500 * 1e6

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	// dp[i][j] 其中j表示还差j堵墙未画
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = max
		}
	}
	dp[0][n] = 0

	for i := 0; i < n; i++ {
		index := (i + 1) % 2
		prevIndex := 1 - index
		count := 1 + time[i]
		copy(dp[index], dp[1-index])
		for j := n; j >= 0; j-- {
			var k int
			if j > count {
				k = j - count
			}
			var temp int
			if dp[prevIndex][k] < dp[prevIndex][j]+cost[i] {
				temp = dp[prevIndex][k]
			} else {
				temp = dp[prevIndex][j] + cost[i]
			}
			if temp < dp[index][k] {
				dp[index][k] = temp
			}
		}
	}
	return dp[n%2][0]
}
