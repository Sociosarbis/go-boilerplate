package nodes2

func dfs1(g [][]int, even, odd []int, c, p int) {
	even[c]++
	for _, next := range g[c] {
		if next != p {
			dfs1(g, even, odd, next, c)
			even[c] += odd[next]
			odd[c] += even[next]
		}
	}
	if p == -1 {
		for _, next := range g[c] {
			dfs2(g, even, odd, next, c)
		}
	}
}

func dfs2(g [][]int, even, odd []int, c, p int) {
	even[c], odd[c] = odd[p], even[p]
	for _, next := range g[c] {
		if next != p {
			dfs2(g, even, odd, next, c)
		}
	}
}

func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {
	n := len(edges1) + 1
	m := len(edges2) + 1
	g1, g2 := make([][]int, n), make([][]int, m)
	for _, edge := range edges1 {
		a, b := edge[0], edge[1]
		g1[a], g1[b] = append(g1[a], b), append(g1[b], a)
	}
	for _, edge := range edges2 {
		a, b := edge[0], edge[1]
		g2[a], g2[b] = append(g2[a], b), append(g2[b], a)
	}
	even1, odd1 := make([]int, n), make([]int, n)
	dfs1(g1, even1, odd1, 0, -1)
	even2, odd2 := make([]int, m), make([]int, m)
	dfs1(g2, even2, odd2, 0, -1)
	var max int
	for i := 0; i < m; i++ {
		if odd2[i] > max {
			max = odd2[i]
		}
	}
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = even1[i] + max
	}
	return ret
}
