package champion

func findChampion(grid [][]int) int {
a:
	for i, row := range grid {
		for j, v := range row {
			if i != j && v == 0 {
				continue a
			}
		}
		return i
	}
	return 0
}
