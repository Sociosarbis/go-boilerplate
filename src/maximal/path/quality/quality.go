package quality

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n := len(values)

	g := make([]map[int]int, n)

	for i := 0; i < n; i++ {
		g[i] = make(map[int]int, 4)
	}

	for _, edge := range edges {
		u, v, time := edge[0], edge[1], edge[2]
		g[u][v] = time
		g[v][u] = time
	}
	var ret int
	dfs(g, values, make([]bool, n), 0, 0, maxTime, 0, &ret)
	return ret
}

func dfs(g []map[int]int, values []int, visited []bool, c, timeCost, maxTime, temp int, ret *int) {
	if timeCost > maxTime {
		return
	}
	firstVisit := !visited[c]
	if firstVisit {
		visited[c] = true
		temp += values[c]
	}
	if c == 0 && temp > *ret {
		*ret = temp
	}

	for next, time := range g[c] {
		dfs(g, values, visited, next, timeCost+time, maxTime, temp, ret)
	}

	if firstVisit {
		visited[c] = false
	}
}
