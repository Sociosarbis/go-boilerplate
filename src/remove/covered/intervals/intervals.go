package intervals

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] > intervals[j][1])
	})
	out := 1
	n := len(intervals)
	end := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][1] > end {
			out++
			end = intervals[i][1]
		}
	}
	return out
}
