package tower

func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := make([][]float64, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]float64, query_row+1)
	}

	dp[0][0] = float64(poured)

	for i := 0; i < query_row; i++ {
		index := i & 1
		nextIndex := 1 - index
		for j := 0; j <= i; j++ {
			if dp[index][j] > 1 {
				v := (dp[index][j] - 1) * 0.5
				dp[nextIndex][j] += v
				dp[nextIndex][j+1] += v
			}
			dp[index][j] = 0
		}
	}
	index := query_row & 1
	if dp[index][query_glass] > 1 {
		return 1
	}
	return dp[index][query_glass]
}
