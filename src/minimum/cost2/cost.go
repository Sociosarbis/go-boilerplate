package cost2

import "sort"

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	sort.Ints(horizontalCut)
	sort.Ints(verticalCut)
	l1 := len(horizontalCut)
	l2 := len(verticalCut)
	c1 := 1
	c2 := 1
	i1 := l1 - 1
	i2 := l2 - 1
	var ret int64
	for i1 >= 0 || i2 >= 0 {
		var isLeft bool
		if i1 >= 0 {
			if i2 >= 0 {
				if horizontalCut[i1] > verticalCut[i2] {
					isLeft = true
				}
			} else {
				isLeft = true
			}
		}
		if isLeft {
			ret += int64(horizontalCut[i1]) * int64(c1)
			i1--
			c2++
		} else {
			ret += int64(verticalCut[i2]) * int64(c2)
			i2--
			c1++
		}
	}
	return ret
}
