package ways

const mask int = 1e9 + 7

func numOfWays(n int) int {
	dp := [2][3][3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != j {
				for k := 0; k < 3; k++ {
					if j != k {
						dp[0][i][j][k] = 1
					}
				}
			}
		}
	}

	for i := 1; i < n; i++ {
		index := i & 1
		prevIndex := 1 - index
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if i != j {
					for k := 0; k < 3; k++ {
						if j != k {
							for ii := 0; ii < 3; ii++ {
								if i != ii {
									for jj := 0; jj < 3; jj++ {
										if jj != j && jj != ii {
											for kk := 0; kk < 3; kk++ {
												if kk != k && kk != jj {
													dp[index][i][j][k] = (dp[index][i][j][k] + dp[prevIndex][ii][jj][kk]) % mask
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	var out int
	index := (n - 1) & 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != j {
				for k := 0; k < 3; k++ {
					if j != k {
						out = (out + dp[index][i][j][k]) % mask
					}
				}
			}
		}
	}
	return out
}
