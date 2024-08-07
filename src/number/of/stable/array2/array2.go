package array2

const mod int = 1e9 + 7

func numberOfStableArrays(zero int, one int, limit int) int {
	dp := make([][][2]int, zero+1)
	prefixSum := make([][][2]int, zero+1)
	for i := 0; i <= zero; i++ {
		dp[i] = make([][2]int, one+1)
		prefixSum[i] = make([][2]int, one+1)
	}

	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			for k := 0; k < 2; k++ {
				n := limit
				if k == 0 {
					if n > i {
						n = i
					}
				} else {
					if n > j {
						n = j
					}
				}
				var temp int
				if k == 0 {
					lower := i - n - 1
					if lower < 0 {
						lower = 0
					}
					temp = (temp + prefixSum[i-1][j][1-k] - prefixSum[lower][j][1-k]) % mod
					if temp < 0 {
						temp += mod
					}
					if i == n && j <= limit {
						temp = (temp + 1) % mod
					}
				} else {
					lower := j - n - 1
					if lower < 0 {
						lower = 0
					}
					temp = (temp + prefixSum[i][j-1][1-k] - prefixSum[i][lower][1-k]) % mod
					if temp < 0 {
						temp += mod
					}
					if j == n && i <= limit {
						temp = (temp + 1) % mod
					}
				}
				dp[i][j][k] = temp
				if k == 0 {
					prefixSum[i][j][k] = (temp + prefixSum[i][j-1][k]) % mod
				} else {
					prefixSum[i][j][k] = (temp + prefixSum[i-1][j][k]) % mod
				}
			}
		}
	}
	return (dp[zero][one][0] + dp[zero][one][1]) % mod
}
