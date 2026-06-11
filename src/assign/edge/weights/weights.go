package weights

const mask int = 1e9 + 7

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func depth(g [][]int, p, c int) int {
	var mx int
	for _, next := range g[c] {
		if next != p {
			mx = max(mx, depth(g, c, next))
		}
	}
	return 1 + mx
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

func assignEdgeWeights(edges [][]int) int {
	n := len(edges) + 1
	g := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g[u], g[v] = append(g[u], v), append(g[v], u)
	}
	d := depth(g, 0, 1)
	return qpow(2, d-2)
}
