package nodes

func countGoodNodes(edges [][]int) int {
	n := len(edges) + 1
	neighbors := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		neighbors[a] = append(neighbors[a], b)
		neighbors[b] = append(neighbors[b], a)
	}
	var ret int
	dfs(neighbors, -1, 0, &ret)
	return ret
}

func dfs(neighbors [][]int, prev, cur int, out *int) int {
	ret := 1
	isSame := true
	prevCount := -1
	for _, next := range neighbors[cur] {
		if next != prev {
			count := dfs(neighbors, cur, next, out)
			if prevCount == -1 {
				prevCount = count
			} else if count != prevCount {
				isSame = false
			}
			ret += count
		}
	}
	if isSame {
		*out++
	}
	return ret
}
