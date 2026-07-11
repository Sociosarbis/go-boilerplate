package components

func dfs(g [][]int, visited []bool, cur, target int) (int, bool) {
	count := 1
	nc := len(g[cur])
	res := nc == target
	for _, next := range g[cur] {
		if !visited[next] {
			visited[next] = true
			subCount, subRes := dfs(g, visited, next, nc)
			res = res && subRes
			count += subCount
		}
	}
	return count, res
}

func countCompleteComponents(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	visited := make([]bool, n)
	var out int
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		count, res := dfs(g, visited, i, len(g[i]))
		if res && count == len(g[i])+1 {
			out++
		}
	}
	return out
}
