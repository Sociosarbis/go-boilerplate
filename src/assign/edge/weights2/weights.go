package weights2

const mask int = 1e9 + 7

func dfs(g, parents [][]int, path, depth []int, p, c, d int) {
	n := len(path)
	for i := 0; 1<<i < n+1; i++ {
		parents[c] = append(parents[c], path[n-(1<<i)])
	}
	path = append(path, c)
	depth[c] = d
	for _, next := range g[c] {
		if next != p {
			dfs(g, parents, path, depth, c, next, d+1)
		}
	}
}

func distance(parents [][]int, depth []int, a, b int) int {
	if depth[a] < depth[b] {
		a, b = b, a
	}
	if depth[a] > depth[b] {
		n := len(parents[a])
		for i := n - 1; i >= 0; i-- {
			if depth[parents[a][i]] >= depth[b] {
				return (1 << i) + distance(parents, depth, parents[a][i], b)
			}
		}
	} else if a != b {
		n := len(parents[a])
		for i := n - 1; i >= 0; i-- {
			if parents[a][i] != parents[b][i] {
				return 1<<(i+1) + distance(parents, depth, parents[a][i], parents[b][i])
			}
		}
		return 2
	}
	return 0
}

func qpow(n int, k int) int {
	var out int = 1
	for k != 0 {
		if k&1 == 1 {
			out = (out * n) % mask
			k--
		} else {
			n = (n * n) % mask
			k >>= 1
		}
	}
	return out
}

func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1
	g := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g[u], g[v] = append(g[u], v), append(g[v], u)
	}
	parents := make([][]int, n+1)
	depth := make([]int, n+1)
	dfs(g, parents, []int{}, depth, 0, 1, 0)
	out := make([]int, len(queries))
	for i, query := range queries {
		u, v := query[0], query[1]
		dist := distance(parents, depth, u, v)
		if dist > 0 {
			out[i] = qpow(2, dist-1)
		}
	}
	return out
}
