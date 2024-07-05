package matrix

func modifiedMatrix(matrix [][]int) [][]int {
	n := len(matrix[0])
	max := make([]int, n)
	for _, row := range matrix {
		for j, num := range row {
			if num > max[j] {
				max[j] = num
			}
		}
	}

	for _, row := range matrix {
		for j, num := range row {
			if num == -1 {
				row[j] = max[j]
			}
		}
	}
	return matrix
}
