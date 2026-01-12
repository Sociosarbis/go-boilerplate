package points

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minTimeToVisitAllPoints(points [][]int) int {
	var out int
	n := len(points)
	for i := 1; i < n; i++ {
		out += max(abs(points[i][0]-points[i-1][0]), abs(points[i][1]-points[i-1][1]))
	}
	return out
}
