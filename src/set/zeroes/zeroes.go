package zeroes

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				if !rows[i] {
					rows[i] = true
				}
				if !cols[j] {
					cols[j] = true
				}
			}
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if rows[i] || cols[j] {
				if matrix[i][j] != 0 {
					matrix[i][j] = 0
				}
			}
		}
	}
}
