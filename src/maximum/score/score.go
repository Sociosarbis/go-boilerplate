package score

import "sort"

func maximumScore(a int, b int, c int) int {
	q := []int{a, b, c}
	sort.Ints(q)
	ret := 0
	for q[1] != 0 {
		q[1]--
		if q[1] < q[0] {
			q[0], q[1] = q[1], q[0]
		}
		q[2]--
		if q[2] < q[1] {
			q[1], q[2] = q[2], q[1]
		}
		ret++
	}
	return ret
}
