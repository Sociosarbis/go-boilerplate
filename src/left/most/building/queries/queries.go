package queries

import (
	"sort"
)

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	n := len(queries)
	sortedIndices := make([]int, n)
	for i, q := range queries {
		sortedIndices[i] = i
		if q[0] > q[1] {
			q[0], q[1] = q[1], q[0]
		}
	}
	sort.Slice(sortedIndices, func(i, j int) bool {
		return queries[sortedIndices[i]][1] > queries[sortedIndices[j]][1]
	})

	hn := len(heights)
	j := hn
	stack := []int{}
	ret := make([]int, n)
	for _, i := range sortedIndices {
		a, b := queries[i][0], queries[i][1]
		if heights[a] < heights[b] || a == b {
			ret[i] = b
		} else {
			for j-1 >= b {
				v := heights[j-1]
				index := sort.Search(len(stack), func(i int) bool {
					return heights[stack[i]] <= v
				})
				if index >= len(stack) {
					stack = append(stack, j-1)
				} else {
					stack = append(stack[:index], j-1)
				}
				j--
			}
			v := heights[a]
			index := sort.Search(len(stack), func(i int) bool {
				return heights[stack[i]] <= v
			})
			if index > 0 {
				ret[i] = stack[index-1]
			} else {
				ret[i] = -1
			}
		}
	}
	return ret
}
