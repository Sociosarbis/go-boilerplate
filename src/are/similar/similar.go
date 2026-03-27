package similar

func normalizeIndex(index, n int) int {
	if index < 0 {
		return (((-index/n)+1)*n + index) % n
	}
	return index % n
}

func areSimilar(mat [][]int, k int) bool {
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i += 2 {
		for j := 0; j < n; j++ {
			index := normalizeIndex(j-k, n)
			if mat[i][j] != mat[i][index] {
				return false
			}
		}
	}
	for i := 1; i < m; i += 2 {
		for j := 0; j < n; j++ {
			index := normalizeIndex(j+k, n)
			if mat[i][j] != mat[i][index] {
				return false
			}
		}
	}
	return true
}
