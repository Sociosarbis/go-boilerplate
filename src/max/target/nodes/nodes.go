package nodes

func dfs(g [][]int, c, p, k int) int {
	if k < 0 {
		return 0
	}
	ret := 1
	if k > 0 {
		for _, next := range g[c] {
			if next != p {
				ret += dfs(g, next, c, k-1)
			}
		}
	}
	return ret
}

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	n, m := len(edges1)+1, len(edges2)+1
	g1, g2 := make([][]int, n), make([][]int, m)
	for _, edge := range edges1 {
		a, b := edge[0], edge[1]
		g1[a] = append(g1[a], b)
		g1[b] = append(g1[b], a)
	}
	for _, edge := range edges2 {
		a, b := edge[0], edge[1]
		g2[a] = append(g2[a], b)
		g2[b] = append(g2[b], a)
	}
	var max int
	for i := 0; i < m; i++ {
		temp := dfs(g2, i, -1, k-1)
		if temp > max {
			max = temp
		}
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = max + dfs(g1, i, -1, k)
	}
	return dp
}
