package subarrays2

import "github.com/boilerplate/src/tree/fenwick"

func countMajoritySubarrays(nums []int, target int) int64 {
	n := len(nums)
	min := -n
	tree := fenwick.NewFenwickTree(2*n + 2)
	var temp int
	var out int64
	tree.Add(-min, 1)
	for _, num := range nums {
		if num == target {
			temp++
		} else {
			temp--
		}
		index := temp - min
		out += int64(tree.Query(index - 1))
		tree.Add(index, 1)
	}
	return out
}
