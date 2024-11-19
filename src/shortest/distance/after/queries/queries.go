package queries

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	dists := make([]int, n)
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		dists[i] = i
		if i+1 < n {
			next[i] = []int{i + 1}
		}
	}
	ret := make([]int, len(queries))
	for i, q := range queries {
		a, b := q[0], q[1]
		temp := dists[a] + 1
		next[a] = append(next[a], b)
		bfs := []int{b}
		m := 1
		for m != 0 {
			for i := 0; i < m; i++ {
				if dists[bfs[i]] > temp {
					dists[bfs[i]] = temp
					bfs = append(bfs, next[bfs[i]]...)
				}
			}
			bfs = bfs[m:]
			temp++
			m = len(bfs)
		}
		ret[i] = dists[n-1]
	}
	return ret
}
