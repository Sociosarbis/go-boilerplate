package count

func maxCount(m int, n int, ops [][]int) int {
	r := []int{m, n}
	for _, op := range ops {
		if op[0] < r[0] {
			r[0] = op[0]
		}
		if op[1] < r[1] {
			r[1] = op[1]
		}
	}
	return r[0] * r[1]
}
