package square

func isOrtho(s1 []int, s2 []int) bool {
	return s1[0]*s2[0]+s1[1]*s2[1] == 0
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	s1 := []int{p2[0] - p1[0], p2[1] - p1[1]}
	s2 := []int{p3[0] - p1[0], p3[1] - p1[1]}
	s3 := []int{p4[0] - p1[0], p4[1] - p1[1]}
	if isOrtho(s1, s2) {
		if isOrtho(s1, s3) {
			return false
		}
		return isOrtho(s3, []int{p3[0] - p2[0], p3[1] - p2[1]}) &&
			isOrtho([]int{p2[0] - p4[0], p2[1] - p4[1]}, []int{p3[0] - p4[0], p3[1] - p4[1]}) &&
			isOrtho(s1, []int{p4[0] - p2[0], p4[1] - p2[1]})
	} else if isOrtho(s1, s3) {
		return isOrtho(s2, []int{p4[0] - p2[0], p4[1] - p2[1]}) &&
			isOrtho([]int{p2[0] - p3[0], p2[1] - p3[1]}, []int{p4[0] - p3[0], p4[1] - p3[1]}) &&
			isOrtho(s1, []int{p3[0] - p2[0], p3[1] - p2[1]})
	} else if isOrtho(s2, s3) {
		return isOrtho(s1, []int{p4[0] - p3[0], p4[1] - p3[1]}) &&
			isOrtho([]int{p3[0] - p2[0], p3[1] - p2[1]}, []int{p4[0] - p2[0], p4[1] - p2[1]}) &&
			isOrtho(s2, []int{p2[0] - p3[0], p2[1] - p3[1]})
	}
	return false
}
