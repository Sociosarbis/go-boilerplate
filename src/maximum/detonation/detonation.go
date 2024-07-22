package detonation

func canEffect(a []int, b []int) bool {
	x := b[0] - a[0]
	y := b[1] - a[1]
	return x*x+y*y <= a[2]*a[2]
}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	g := make([][]int, n)

	for i, bomb := range bombs {
		g[i] = []int{}
		for j := i + 1; j < n; j++ {
			if canEffect(bomb, bombs[j]) {
				g[i] = append(g[i], j)
			}
		}
	}

	var ret int
	for i := 0; i < n; i++ {
		res := dfs(g, i, make([]bool, n))
		if res > ret {
			ret = res
		}
	}
	return ret
}

func dfs(g [][]int, i int, visited []bool) int {
	var ret int
	for _, next := range g[i] {
		if !visited[next] {
			visited[next] = true
			ret++
			ret += dfs(g, next, visited)
		}
	}
	return ret
}
