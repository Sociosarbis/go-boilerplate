package methods

func remainingMethods(n int, k int, invocations [][]int) []int {
	counter := make([]int, n)
	g := make([][]int, n)
	for _, invocation := range invocations {
		a, b := invocation[0], invocation[1]
		counter[b]++
		g[a] = append(g[a], b)
	}
	visited := make([]bool, n)
	bfs := []int{k}
	group := make([]int, 0, n)
	group = append(group, k)
	visited[k] = true
	l := len(bfs)
	for l != 0 {
		for i := 0; i < l; i++ {
			for _, next := range g[bfs[i]] {
				if !visited[next] {
					visited[next] = true
					bfs = append(bfs, next)
					group = append(group, next)
				}
				counter[next]--
			}
		}
		bfs = bfs[l:]
		l = len(bfs)
	}
	for _, method := range group {
		if counter[method] != 0 {
			out := make([]int, n)
			for i := 0; i < n; i++ {
				out[i] = i
			}
			return out
		}
	}
	out := make([]int, 0, n-len(group))
	for i := 0; i < n; i++ {
		if !visited[i] {
			out = append(out, i)
		}
	}
	return out
}
