package trapezoids

import (
	"sort"
)

const mask int64 = 1e9 + 7

func countTrapezoids(points [][]int) int {
	n := len(points)
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	counts := make([]int, 0, n)
	temp := 1
	for i := 1; i < n; i++ {
		if points[i][1] == points[i-1][1] {
			temp++
		} else {
			if temp > 1 {
				counts = append(counts, temp-1)
			}
			temp = 1
		}
	}
	if temp > 1 {
		counts = append(counts, temp-1)
	}
	n = len(counts)
	var sum int64
	for i := 0; i < n; i++ {
		sum = (sum + (int64(counts[i])+1)*int64(counts[i])/2) % mask
	}
	var out int
	for i := 0; i < n; i++ {
		a := (int64(counts[i]) + 1) * int64(counts[i]) / 2
		if sum < a {
			sum = sum + mask - a
		} else {
			sum = sum - a
		}
		out = (out + int(a*sum%mask)) % int(mask)
	}
	return out
}
