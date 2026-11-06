package servers

import "sort"

func countServers(n int, logs [][]int, x int, queries []int) []int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})
	m := len(queries)
	qi := make([]int, len(queries))
	for i := range qi {
		qi[i] = i
	}
	sort.Slice(qi, func(i, j int) bool {
		return queries[qi[i]] < queries[qi[j]]
	})
	var temp int
	out := make([]int, m)
	counter := make([]int, n+1)
	m = len(logs)
	var l, r int
	for i, index := range qi {
		query := queries[index]
		for r < m && logs[r][1] <= query {
			counter[logs[r][0]]++
			if counter[logs[r][0]] == 1 {
				temp++
			}
			r++
		}
		for l < r && logs[l][1] < query-x {
			counter[logs[l][0]]--
			if counter[logs[l][0]] == 0 {
				temp--
			}
			l++
		}
		out[i] = n - temp
	}
	return out
}
