package wood

func sellingWood(m int, n int, prices [][]int) int64 {
	dp := make([][]int64, m+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}
	pMap := make([][]int, m+1)
	for i := range pMap {
		pMap[i] = make([]int, n+1)
	}

	for _, price := range prices {
		w, h, p := price[0], price[1], price[2]
		pMap[w][h] = p
	}

	var ret int64
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			var res int64 = int64(pMap[i][j])
			for k := 0; k < i; k++ {
				temp := dp[k][j] + dp[i-k][j]
				if temp > res {
					res = temp
				}
			}

			for k := 0; k < j; k++ {
				temp := dp[i][k] + dp[i][j-k]
				if temp > res {
					res = temp
				}
			}
			dp[i][j] = res
			if res > ret {
				ret = res
			}
		}
	}
	return ret
}
