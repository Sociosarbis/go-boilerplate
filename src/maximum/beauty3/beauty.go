package beauty3

import "sort"

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	n := len(items)
	var j int
	for i := 1; i < n; i++ {
		top := items[j]
		if top[0] == items[i][0] {
			if items[i][1] > top[1] {
				top[1] = items[i][1]
			}
		} else if items[i][1] > top[1] {
			j++
			items[j] = items[i]
		}
	}
	items = items[:j+1]
	n = len(items)
	ret := make([]int, len(queries))
	for i, q := range queries {
		var l int
		r := n - 1
		for l <= r {
			mid := (l + r) >> 1
			if items[mid][0] > q {
				r = mid - 1
			} else if mid+1 < n && items[mid+1][0] <= q {
				l = mid + 1
			} else {
				ret[i] = items[mid][1]
				break
			}
		}
	}
	return ret
}
