package pairs2

func isUpperLeft(p1, p2 []int) bool {
	return p1[0] <= p2[0] && p1[1] >= p2[1]
}

func isIncluded(p1, p2, p3 []int) bool {
	return p3[0] >= p1[0] && p3[1] <= p1[1] && p3[0] <= p2[0] && p3[1] >= p2[1]
}

func numberOfPairs(points [][]int) int {
	n := len(points)
	var out int
	for i, point := range points {
	loop:
		for j := i + 1; j < n; j++ {
			a, b := point, points[j]
			if isUpperLeft(a, b) || isUpperLeft(b, a) {
				if isUpperLeft(b, a) {
					a, b = b, a
				}
				for k, point := range points {
					if k != i && k != j {
						if isIncluded(a, b, point) {
							continue loop
						}
					}
				}
				out++
			}
		}
	}
	return out
}
