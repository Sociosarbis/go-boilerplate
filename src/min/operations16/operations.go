package operations16

import "sort"

func minOperations(grid [][]int, x int) int {
	m, n := len(grid), len(grid[0])
	nums := make([]int, 0, m*n)
	first := grid[0][0]
	for _, row := range grid {
		for _, num := range row {
			if (num-first)%x != 0 {
				return -1
			}
			nums = append(nums, num)
		}
	}
	sort.Ints(nums)
	mid := nums[(m*n)/2]
	var out int
	for _, num := range nums {
		if num < mid {
			out += (mid - num) / x
		} else if num > mid {
			out += (num - mid) / x
		}
	}
	return out
}
