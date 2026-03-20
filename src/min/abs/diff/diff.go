package diff

import "sort"

func minAbsDiff(grid [][]int, k int) [][]int {
	m, n := len(grid)-k+1, len(grid[0])-k+1
	out := make([][]int, m)
	l := k * k
	temp := make([]int, l)
	for i := 0; i < m; i++ {
		out[i] = make([]int, n)
		for j := 0; j < n; j++ {
			var index int
			for x := 0; x < k; x++ {
				for y := 0; y < k; y++ {
					temp[index] = grid[i+x][j+y]
					index++
				}
			}
			sort.Ints(temp)
			min := temp[l-1] - temp[0]
			for index = 1; index < l; index++ {
				if temp[index] != temp[index-1] && temp[index]-temp[index-1] < min {
					min = temp[index] - temp[index-1]
				}
			}
			out[i][j] = min
		}
	}
	return out
}
