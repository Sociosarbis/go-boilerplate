package distance2

import "sort"

func maxDistance(arrays [][]int) int {
	min, max := [][2]int{{0, arrays[0][0]}, {1, arrays[1][0]}}, [][2]int{{0, arrays[0][len(arrays[0])-1]}, {1, arrays[1][len(arrays[1])-1]}}
	sort.Slice(min, func(i, j int) bool {
		return min[i][1] < min[j][1]
	})
	sort.Slice(max, func(i, j int) bool {
		return max[i][1] > max[j][1]
	})
	n := len(arrays)
	for i := 2; i < n; i++ {
		l, r := arrays[i][0], arrays[i][len(arrays[i])-1]
		if l < min[0][1] {
			min[0], min[1] = [2]int{i, l}, min[0]
		} else if l < min[1][1] {
			min[1] = [2]int{i, l}
		}

		if r > max[0][1] {
			max[0], max[1] = [2]int{i, r}, max[0]
		} else if r > max[1][1] {
			max[1] = [2]int{i, r}
		}
	}
	if max[0][0] != min[0][0] {
		return max[0][1] - min[0][1]
	}
	a, b := max[0][1]-min[1][1], max[1][1]-min[0][1]
	if a > b {
		return a
	}
	return b
}
