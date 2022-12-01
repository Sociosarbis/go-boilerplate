package point

func nearestValidPoint(x int, y int, points [][]int) int {
	ret := -1
	temp := -1
	dist := 0
	for i, point := range points {
		if point[0] == x {
			dist = point[1] - y
		} else if point[1] == y {
			dist = point[0] - x
		} else {
			continue
		}
		if dist < 0 {
			dist = -dist
		}
		if temp == -1 || dist < temp {
			temp = dist
			ret = i
		}
	}
	return ret
}
