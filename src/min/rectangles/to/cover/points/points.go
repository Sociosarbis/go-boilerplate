package points

import "sort"

func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	ret := 1
	start := points[0][0]
	n := len(points)
	for i := 1; i < n; i++ {
		if points[i][0]-start > w {
			start = points[i][0]
			ret++
		}
	}
	return ret
}
