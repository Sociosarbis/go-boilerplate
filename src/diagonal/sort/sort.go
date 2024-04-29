package sort

import "sort"

func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	for j := 0; j < n; j++ {
		var l int
		if n-j > m {
			l = m
		} else {
			l = n - j
		}
		row := make([]int, l)
		for k := 0; k < l; k++ {
			row[k] = mat[k][j+k]
		}
		sort.Ints(row)
		for k := 0; k < l; k++ {
			mat[k][j+k] = row[k]
		}
	}

	for i := 1; i < m; i++ {
		var l int
		if m-i > n {
			l = n
		} else {
			l = m - i
		}
		row := make([]int, l)
		for k := 0; k < l; k++ {
			row[k] = mat[i+k][k]
		}
		sort.Ints(row)
		for k := 0; k < l; k++ {
			mat[i+k][k] = row[k]
		}
	}
	return mat
}
