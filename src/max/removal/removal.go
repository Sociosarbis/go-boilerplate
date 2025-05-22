package removal

import (
	"container/heap"
	"sort"

	h "github.com/boilerplate/src/heap"
)

func maxRemoval(nums []int, queries [][]int) int {
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})

	hp := h.New[[]int]([][]int{}, func(a, b []int) bool {
		return a[1] > b[1]
	})

	var ret int
	var temp int
	var i int
	n := len(queries)
	m := len(nums)
	dp := make([]int, m)
	for j, num := range nums {
		temp += dp[j]
		for i < n && j >= queries[i][0] {
			if queries[i][1] >= j {
				heap.Push(&hp, queries[i])
			}
			i++
		}
		for num > temp && hp.Len() > 0 {
			top := heap.Pop(&hp).([]int)
			if top[1] >= j {
				temp++
				ret++
				if top[1]+1 < m {
					dp[top[1]+1]--
				}
			}
		}
		if num > temp {
			return -1
		}
	}
	return n - ret
}
