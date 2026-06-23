package arrays

const mod int = 1e9 + 7

func getSum(prefix []int, l, r int) int {
	out := prefix[r+1] - prefix[l]
	if out < 0 {
		return out + mod
	}
	return out
}

func zigZagArrays(n int, l int, r int) int {
	dp := [2][]int{}
	m := r - l + 1
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			dp[1][j]++
		}
		for j := 0; j < i; j++ {
			dp[0][j]++
		}
	}
	prefix := [2][]int{}
	for i := 0; i < 2; i++ {
		prefix[i] = make([]int, m+1)
		for j := 0; j < m; j++ {
			prefix[i][j+1] = (prefix[i][j] + dp[i][j]) % mod
		}
	}

	for i := 2; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[0][j], dp[1][j] = 0, 0
		}
		for j := 0; j < m; j++ {
			dp[1][j] = getSum(prefix[0], 0, j-1)
			dp[0][j] = getSum(prefix[1], j+1, m-1)
		}
		for i := 0; i < 2; i++ {
			prefix[i][0] = 0
			for j := 0; j < m; j++ {
				prefix[i][j+1] = (prefix[i][j] + dp[i][j]) % mod
			}
		}
	}
	var out int
	for i := 0; i < m; i++ {
		out = (out + (dp[0][i]+dp[1][i])%mod) % mod
	}
	return out
}
