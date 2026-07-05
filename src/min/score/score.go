package score

import "fmt"

func hash(a, b int) string {
	if a > b {
		a, b = b, a
	}
	return fmt.Sprintf("%v_%v", a, b)
}

type empty struct{}

func dfs(g [][][2]int, visited map[string]empty, cur int) int {
	out := 10000
	for _, next := range g[cur] {
		if out > next[1] {
			out = next[1]
		}
		h := hash(cur, next[0])
		if _, ok := visited[h]; ok {
			continue
		}
		visited[h] = empty{}
		temp := dfs(g, visited, next[0])
		if temp < out {
			out = temp
		}
	}
	return out
}

func minScore(n int, roads [][]int) int {
	g := make([][][2]int, n)
	for _, road := range roads {
		a, b, score := road[0], road[1], road[2]
		g[a-1] = append(g[a-1], [2]int{b - 1, score})
		g[b-1] = append(g[b-1], [2]int{a - 1, score})
	}
	visited := make(map[string]empty, len(roads))
	res := dfs(g, visited, 0)

	return res
}
