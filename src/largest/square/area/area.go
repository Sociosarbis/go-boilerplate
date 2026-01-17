package area

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)
	var maxSide int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x := min(topRight[i][0], topRight[j][0]) - max(bottomLeft[i][0], bottomLeft[j][0])
			if x <= 0 {
				continue
			}
			y := min(topRight[i][1], topRight[j][1]) - max(bottomLeft[i][1], bottomLeft[j][1])
			if y <= 0 {
				continue
			}
			side := min(x, y)
			if side > maxSide {
				maxSide = side
			}
		}
	}
	return int64(maxSide) * int64(maxSide)
}
