package queries

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	var i int
	groups := make([]int, n)
	for j := 1; j < n; j++ {
		if nums[j]-nums[j-1] > maxDiff {
			i++
		}
		groups[j] = i
	}
	out := make([]bool, len(queries))
	for i, query := range queries {
		out[i] = groups[query[0]] == groups[query[1]]
	}
	return out
}
