package queries

func rangeAddQueries(n int, queries [][]int) [][]int {
	out := make([][]int, n)
	for i := 0; i < n; i++ {
		out[i] = make([]int, n)
	}
	for _, query := range queries {
		r1, c1, r2, c2 := query[0], query[1], query[2], query[3]
		c2Next := c2 + 1
		for i := r1; i <= r2; i++ {
			out[i][c1]++
			if c2Next < n {
				out[i][c2Next]--
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			out[i][j] += out[i][j-1]
		}
	}
	return out
}
