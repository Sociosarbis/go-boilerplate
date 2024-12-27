package element

import "sort"

func occurrencesOfElement(nums []int, queries []int, x int) []int {
	n := len(queries)
	ret := make([]int, n)
	qs := make([][2]int, n)
	for i := 0; i < n; i++ {
		qs[i] = [2]int{queries[i], i}
	}
	sort.Slice(qs, func(i, j int) bool {
		return qs[i][0] < qs[j][0]
	})

	var c int
	var j int
	for k, num := range nums {
		if num == x {
			c++
		}
		for j < n && c > qs[j][0] {
			j++
		}
		for j < n && c == qs[j][0] {
			ret[qs[j][1]] = k
			j++
		}
	}
	for j < n {
		ret[qs[j][1]] = -1
		j++
	}
	return ret
}
