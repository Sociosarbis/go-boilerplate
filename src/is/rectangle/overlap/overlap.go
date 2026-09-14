package overlap

func isOverlap(l1, r1, l2, r2 int) bool {
	return !(l1 >= r2 || r1 <= l2)
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return isOverlap(rec1[0], rec1[2], rec2[0], rec2[2]) && isOverlap(rec1[1], rec1[3], rec2[1], rec2[3])
}
