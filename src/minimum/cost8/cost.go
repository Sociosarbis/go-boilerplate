package cost8

import "sort"

func minimumCost(cost []int) int {
	sort.Ints(cost)
	n := len(cost)
	var out int
	for i := n - 1; i >= 0; i -= 3 {
		for j := 0; j < 2 && j <= i; j++ {
			out += cost[i-j]
		}
	}
	return out
}
