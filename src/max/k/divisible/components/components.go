package components

func dfs(g [][]int, values []int, prev, cur, k int) (total int, sum int) {
	total = values[cur]
	for _, next := range g[cur] {
		if next != prev {
			t, s := dfs(g, values, cur, next, k)
			sum += s
			total += t
		}
	}
	if total%k == 0 {
		sum++
		total = 0
	}
	return total, sum
}

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	g := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	_, sum := dfs(g, values, -1, 0, k)
	return sum
}
