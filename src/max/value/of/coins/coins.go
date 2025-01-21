package coins

func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	dp := [2][]int{
		make([]int, k+1),
		make([]int, k+1),
	}
	m := len(piles[0])
	var temp int
	for i := k - 1; i >= 0; i-- {
		if k-i <= m {
			temp += piles[0][k-i-1]
			dp[0][i] = temp
		} else {
			dp[0][i] = -1
		}
	}
	var index int
	var prevIndex int
	for i := 1; i < n; i++ {
		m = len(piles[i])
		index = i & 1
		prevIndex = 1 - index
		copy(dp[index], dp[prevIndex])
		temp = 0
		for j := 0; j < m; j++ {
			temp += piles[i][j]
			for ii := k; ii-j-1 >= 0; ii-- {
				if dp[prevIndex][ii] != -1 {
					jj := ii - j - 1
					a := dp[prevIndex][ii] + temp
					if a > dp[index][jj] {
						dp[index][jj] = a
					}
				}
			}
		}
	}
	index = (n - 1) & 1
	return dp[index][0]
}
