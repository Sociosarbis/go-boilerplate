package reorder

func minReorder(n int, connections [][]int) int {
	out := make([][]int, n)
	in := make([][]int, n)
	for _, connection := range connections {
		out[connection[0]] = append(out[connection[0]], connection[1])
		in[connection[1]] = append(in[connection[1]], connection[0])
	}
	visited := make([]bool, n)
	visited[0] = true
	return dfs(out, in, 0, visited)
}

func dfs(out [][]int, in [][]int, i int, visited []bool) int {
	ret := 0
	for _, next := range in[i] {
		if !visited[next] {
			visited[next] = true
			ret += dfs(out, in, next, visited)
		}
	}
	for _, next := range out[i] {
		if !visited[next] {
			visited[next] = true
			ret += 1
			ret += dfs(out, in, next, visited)
		}
	}
	return ret
}
