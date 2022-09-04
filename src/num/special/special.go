package special

func numSpecial(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	rowCounts := make([]int, m)
	colCounts := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				rowCounts[i]++
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[j][i] == 1 {
				colCounts[i]++
			}
		}
	}
	ret := 0
	for i := 0; i < m; i++ {
		if rowCounts[i] == 1 {
			for j := 0; j < n; j++ {
				if mat[i][j] == 1 && colCounts[j] == 1 {
					ret++
				}
			}
		}
	}
	return ret
}
