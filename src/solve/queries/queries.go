package queries

func solveQueries(nums []int, queries []int) []int {
	n := len(nums)
	dists := make([]int, n)
	for i := 0; i < n; i++ {
		dists[i] = n
	}
	numToIndex := make(map[int]int, n)
	m := n * 2
	for i := 0; i < m; i++ {
		index := i % n
		if j, ok := numToIndex[nums[index]]; ok {
			dist := i - j
			if dist < dists[index] {
				dists[index] = dist
			}
			prevIndex := j % n
			if dist < dists[prevIndex] {
				dists[prevIndex] = dist
			}
		}
		numToIndex[nums[index]] = i
	}
	out := make([]int, len(queries))

	for i, query := range queries {
		out[i] = dists[query]
	}
	return out
}
