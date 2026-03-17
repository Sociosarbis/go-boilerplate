package submatrix

import "sort"

func largestSubmatrix(matrix [][]int) int {
	n := len(matrix[0])
	row := make([]int, n)
	temp := make([]int, n)
	copy(row, matrix[0])
	copy(temp, row)
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] > temp[j]
	})
	var out int
	for i, num := range temp {
		if num == 0 {
			break
		}
		out = i + 1
	}
	m := len(matrix)
	for i := 1; i < m; i++ {
		for j, num := range matrix[i] {
			if num == 0 {
				row[j] = 0
			} else {
				row[j]++
			}
		}
		copy(temp, row)
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] > temp[j]
		})
		min := temp[0]
		for i, num := range temp {
			if num == 0 {
				break
			}
			if num < min {
				min = num
			}
			if (i+1)*min > out {
				out = (i + 1) * min
			}
		}
	}
	return out
}
