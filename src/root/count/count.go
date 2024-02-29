package count

type empty struct{}

func rootCount(edges [][]int, guesses [][]int, k int) int {
	n := len(edges) + 1
	g := make([][]int, n)

	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	guessMap := map[int]map[int]empty{}
	for _, guess := range guesses {
		p, c := guess[0], guess[1]
		set(guessMap, p, c)
	}

	ret := 0
	count := 0
	visited := map[int]map[int]empty{}
	dfs1(g, 0, -1, &count, guessMap, visited)
	dfs2(g, 0, -1, &count, guessMap, visited, k, &ret)
	return ret
}

func set(m map[int]map[int]empty, p, c int) {
	if _, ok := m[p]; !ok {
		m[p] = map[int]empty{}
	}
	m[p][c] = empty{}
}

func has(m1 map[int]map[int]empty, p, c int) bool {
	if m2, ok := m1[p]; ok {
		if _, ok = m2[c]; ok {
			return true
		}
	}
	return false
}

func dfs1(g [][]int, i int, p int, count *int, guessMap map[int]map[int]empty, visited map[int]map[int]empty) {
	if has(guessMap, p, i) {
		set(visited, p, i)
		*count++
	}

	for _, next := range g[i] {
		if next != p {
			dfs1(g, next, i, count, guessMap, visited)
		}
	}
}

func dfs2(g [][]int, i int, p int, count *int, guessMap map[int]map[int]empty, visited map[int]map[int]empty, k int, out *int) {
	if *count >= k {
		*out++
	}

	for _, next := range g[i] {
		if next != p {
			if has(visited, i, next) {
				*count--
				delete(visited[i], next)
			}
			if has(guessMap, next, i) {
				*count++
				set(visited, next, i)
			}
			dfs2(g, next, i, count, guessMap, visited, k, out)
			if has(guessMap, i, next) {
				*count++
				set(visited, i, next)
			}
			if has(visited, next, i) {
				*count--
				delete(visited[next], i)
			}
		}
	}
}
