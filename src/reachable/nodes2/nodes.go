package nodes2

type empty struct{}

func reachableNodes(n int, edges [][]int, restricted []int) int {
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	rMap := make(map[int]empty, len(restricted))
	for _, num := range restricted {
		rMap[num] = empty{}
	}
	return dfs(g, 0, -1, rMap)
}

func dfs(g [][]int, i, p int, restricted map[int]empty) int {
	ret := 1
	for _, next := range g[i] {
		if next != p {
			if _, ok := restricted[next]; !ok {
				ret += dfs(g, next, i, restricted)
			}
		}
	}
	return ret
}
