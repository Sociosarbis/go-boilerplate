package sum

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func maxMatrixSum(matrix [][]int) int64 {
	var minAbs int = 1e5
	var out int64
	var odd bool
	for _, row := range matrix {
		for _, num := range row {
			if num < 0 {
				odd = !odd
			}
			num = abs(num)
			minAbs = min(minAbs, num)
			out += int64(num)
		}
	}
	if odd {
		return out - int64(minAbs)*2
	}
	return out
}
