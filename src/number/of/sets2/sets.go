package sets2

const mask int = 1e9 + 7

func numberOfSets(n int, k int) int {
	dp := make([]int, 1+k)
	p1, p2 := make([]int, 1+k), make([]int, 1+k)
	p1[0], p2[0] = 1, 1
	var out int
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[j] = p1[j-1]
			p2[j-1] = (p2[j-1] + dp[j-1]) % mask
			p1[j-1] = (p1[j-1] + p2[j-1]) % mask
		}
		p2[k] = (p2[k] + dp[k]) % mask
		p1[k] = (p1[k] + p2[k]) % mask
		out = (out + dp[k]) % mask
	}
	return out
}
