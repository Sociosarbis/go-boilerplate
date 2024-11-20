package queries2

// 题目保证两条路不会有交叠，但可以包含
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	nxt := make([]int, n-1)
	for i := 1; i < n; i++ {
		nxt[i-1] = i
	}
	cnt := n - 1
	ret := make([]int, len(queries))
	for i, q := range queries {
		u, v := q[0], q[1]
		if nxt[u] > 0 && nxt[u] < v {
			cur := nxt[u]
			for cur < v {
				nxt[cur], cur = 0, nxt[cur]
				cnt--
			}
			nxt[u] = v
		}
		ret[i] = cnt
	}
	return ret
}
