package width

func getWidth(num int) int {
	var ret int
	if num <= 0 {
		ret = 1
	}
	for num != 0 {
		num /= 10
		ret++
	}
	return ret
}

func findColumnWidth(grid [][]int) []int {
	n := len(grid[0])
	ret := make([]int, n)
	for _, row := range grid {
		for i, num := range row {
			width := getWidth(num)
			if width > ret[i] {
				ret[i] = width
			}
		}
	}
	return ret
}
