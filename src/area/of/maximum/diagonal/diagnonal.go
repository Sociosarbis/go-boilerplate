package diagonal

func areaOfMaxDiagonal(dimensions [][]int) int {
	var diagnonal, area int
	for _, dimension := range dimensions {
		d, a := dimension[0]*dimension[0]+dimension[1]*dimension[1], dimension[0]*dimension[1]
		if d > diagnonal || (d == diagnonal && a > area) {
			diagnonal, area = d, a
		}
	}
	return area
}
