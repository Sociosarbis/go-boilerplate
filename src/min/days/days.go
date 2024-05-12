package days

func minDays(n int) int {
	dp := map[int]int{
		1: 1,
	}
	return dfs(n, dp)
}

func dfs(n int, dp map[int]int) int {
	if n == 0 {
		return 0
	}
	if c, ok := dp[n]; ok {
		return c
	}
	a := n%2 + dfs(n/2, dp)
	b := n%3 + dfs(n/3, dp)
	var temp int
	if a < b {
		temp = 1 + a
	} else {
		temp = 1 + b
	}
	dp[n] = temp
	return temp
}
