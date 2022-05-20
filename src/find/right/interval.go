package interval

import (
	"sort"
)

type item [][]int

func (this item) Len() int {
	return len(this)
}

func (this item) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this item) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}

func findRightInterval(intervals [][]int) []int {
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}
	sort.Sort(item(intervals))

	ret := make([]int, len(intervals))

	for i, interval := range intervals {
		left := i
		right := len(intervals) - 1
		target := interval[1]
		ret[intervals[i][2]] = -1
		for left <= right {
			mid := (left + right) >> 1

			if intervals[mid][0] < target {
				left = mid + 1
			} else if intervals[mid][0] > target {
				if mid-1 < left || intervals[mid-1][0] < target {
					ret[intervals[i][2]] = intervals[mid][2]
					break
				} else {
					right = mid - 1
				}
			} else {
				ret[intervals[i][2]] = intervals[mid][2]
				break
			}
		}
	}
	return ret
}
