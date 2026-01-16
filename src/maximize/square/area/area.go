package area

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	hFences, vFences = append(hFences, 1, m), append(vFences, 1, n)
	sort.Ints(hFences)
	sort.Ints(vFences)
	m, n = len(hFences), len(vFences)
	count := max(m, n)
	visited := make(map[int]bool, count*count)
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			diff := hFences[j] - hFences[i]
			if _, ok := visited[diff]; !ok {
				visited[diff] = true
			}
		}
	}
	out := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diff := vFences[j] - vFences[i]
			if _, ok := visited[diff]; ok {
				if diff > out {
					out = diff
				}
			}
		}
	}
	if out == -1 {
		return out
	}
	return int((int64(out) * int64(out)) % (1e9 + 7))
}
