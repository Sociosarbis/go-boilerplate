package rotation

func isEqual(mat, target [][]int, transform func(i, j int) (int, int)) bool {
	for i, row := range mat {
		for j, cell := range row {
			targetI, targetJ := transform(i, j)
			if cell != target[targetI][targetJ] {
				return false
			}
		}
	}
	return true
}

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	return isEqual(mat, target, func(i, j int) (int, int) {
		return i, j
	}) || isEqual(mat, target, func(i, j int) (int, int) {
		return j, n - i - 1
	}) || isEqual(mat, target, func(i, j int) (int, int) {
		return n - i - 1, n - j - 1
	}) || isEqual(mat, target, func(i, j int) (int, int) {
		return n - j - 1, i
	})
}
