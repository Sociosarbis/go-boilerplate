package points

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const maxLevel = 14
const maxN = maxLevel + 1

func dfs(g [][]int, coins []int, dp [][maxN]int, p, c, k int) {
	for _, next := range g[c] {
		if next != p {
			dfs(g, coins, dp, c, next, k)
		}
	}
	v := coins[c]
	for i := 0; i+1 < maxN; i++ {
		temp := v >> i
		a, b := temp-k, temp>>1
		for _, next := range g[c] {
			if next != p {
				a += dp[next][i]
				b += dp[next][i+1]
			}
		}
		dp[c][i] = max(a, b)
	}
}

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	dp := make([][maxN]int, n)
	dfs(g, coins, dp, -1, 0, k)
	return dp[0][0]
}
