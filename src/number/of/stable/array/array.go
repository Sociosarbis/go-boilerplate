package array

const mod int = 1e9 + 7

func numberOfStableArrays(zero int, one int, limit int) int {
	dp := make([][][2][]int, zero+1)
	for i := 0; i <= zero; i++ {
		dp[i] = make([][2][]int, one+1)
		for j := 0; j <= one; j++ {
			dp[i][j] = [2][]int{make([]int, limit+1), make([]int, limit+1)}
		}
	}

	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			var count int
			if i+j < limit {
				count = i + j
			} else {
				count = limit
			}
			for k := 0; k < 2; k++ {
				n := count
				if k == 0 {
					if n > i {
						n = i
					}
				} else {
					if n > j {
						n = j
					}
				}
				ii := i
				jj := j
				if k == 0 {
					ii -= 1
				} else {
					jj -= 1
				}
				for l := 1; l <= n; l++ {
					temp := dp[i][j][k][l]
					if l == 1 {
						if (ii == 0 && jj <= limit) || (jj == 0 && ii <= limit) {
							temp = (temp + 1) % mod
						} else {
							for ll := 1; ll <= count; ll++ {
								temp = (temp + dp[ii][jj][1-k][ll]) % mod
							}
						}
					}
					dp[i][j][k][l] = (temp + dp[ii][jj][k][l-1]) % mod
				}
			}
		}
	}
	var ret int
	for i := 0; i < 2; i++ {
		for j := 1; j <= limit; j++ {
			ret = (ret + dp[zero][one][i][j]) % mod
		}
	}
	return ret
}
