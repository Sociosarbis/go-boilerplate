package rotate

func rotate(matrix [][]int) {
	n := len(matrix)
	tmp1 := make([]int, n)
	for j := 0; j < n; j++ {
		c := 0
		for i := n - 1; i >= j; i-- {
			tmp1[c] = matrix[i][j]
			c++
		}
		for i := j - 1; i >= 0; i-- {
			tmp1[c] = matrix[j][i]
			c++
		}
		for i := j + 1; i < n; i++ {
			matrix[i][j] = matrix[j][i]
		}
		copy(matrix[j][:], tmp1)
	}
}
