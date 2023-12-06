package price

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	graph := make([][]int, n)
	counter := make([]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	for _, trip := range trips {
		visited := make([]bool, n)
		visited[trip[0]] = true
		dfs1(graph, trip[0], trip[1], visited)
		for i, ok := range visited {
			if ok {
				counter[i] += 1
			}
		}
	}
	visited := make([]bool, n)
	visited[0] = true
	ret := dfs2(graph, 0, counter, price, visited)
	full_cost := 0
	for i, count := range counter {
		full_cost += price[i] * count
	}
	if ret[0] > ret[1] {
		return full_cost - ret[0]
	}
	return full_cost - ret[1]
}

func dfs1(graph [][]int, i int, end int, visited []bool) bool {
	if end == i {
		return true
	}
	for _, next := range graph[i] {
		if !visited[next] {
			visited[next] = true
			res := dfs1(graph, next, end, visited)
			if res {
				return res
			}
			visited[next] = false
		}
	}
	return false
}

func dfs2(graph [][]int, i int, counter []int, price []int, visited []bool) []int {
	ret := []int{
		0,
		counter[i] * price[i] / 2,
	}
	for _, next := range graph[i] {
		if !visited[next] {
			visited[next] = true
			res := dfs2(graph, next, counter, price, visited)
			if res[0] > res[1] {
				ret[0] += res[0]
			} else {
				ret[0] += res[1]
			}
			ret[1] += res[0]
		}
	}
	return ret
}
