package servers

type endPoint struct {
	v int
	d int
}

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	g := make([][]endPoint, n)
	for _, edge := range edges {
		a, b, w := edge[0], edge[1], edge[2]
		g[a] = append(g[a], endPoint{b, w})
		g[b] = append(g[b], endPoint{a, w})
	}
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		n := len(g[i])
		groups := make([]int, n)
		var sum int
		for j, end := range g[i] {
			groups[j] = dfs(g, i, end.v, end.d, signalSpeed)
			sum += groups[j]
		}
		var temp int
		for j := 0; j < n; j++ {
			temp += (sum - groups[j]) * groups[j]
		}
		ret[i] = temp / 2
	}
	return ret
}

func dfs(g [][]endPoint, p int, c int, d int, signalSpeed int) int {
	var ret int
	if d%signalSpeed == 0 {
		ret = 1
	}
	for _, end := range g[c] {
		if end.v != p {
			ret += dfs(g, c, end.v, d+end.d, signalSpeed)
		}
	}
	return ret
}
