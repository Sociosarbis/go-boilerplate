package pairs3

import "sort"

func numberOfPairs(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || (points[i][0] == points[j][0] && points[i][1] > points[j][1])
	})

	var out int
	n := len(points)
	for i, a := range points {
		var max int = -1e9 - 1
		for j := i + 1; j < n; j++ {
			if points[j][1] <= a[1] {
				if max < points[j][1] {
					max = points[j][1]
					out++
				}
			}
		}
	}
	return out
}
