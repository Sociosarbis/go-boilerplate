package points

func countPoints(points [][]int, queries [][]int) []int {
	ret := make([]int, len(queries))
	for i, query := range queries {
		rSquare := query[2] * query[2]
		for _, point := range points {
			x := point[0] - query[0]
			y := point[1] - query[1]
			if x*x+y*y <= rSquare {
				ret[i]++
			}
		}
	}
	return ret
}
