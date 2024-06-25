package matrix

import "sort"

func getBits(row []int) int {
	var mask int
	for i, num := range row {
		if num == 1 {
			mask |= 1 << i
		}
	}
	return mask
}

func goodSubsetofBinaryMatrix(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	count := 1 << n
	groups := make([][]int, count)

	for i := m - 1; i >= 0; i-- {
		bits := getBits(grid[i])
		groups[bits] = append(groups[bits], i)
	}
	ret := groups[0]
	for i := 1; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if i&j == 0 {
				var c3 int
				c1 := len(groups[i])
				c2 := len(groups[j])
				if c1 < c2 {
					c3 = c1
				} else {
					c3 = c2
				}
				if c3 != 0 {
					ret = append(append(ret, groups[i][c1-c3:]...), groups[j][c2-c3:]...)
					groups[i] = groups[i][:c1-c3]
					groups[j] = groups[j][:c2-c3]
				}
			}
		}
	}
	sort.Ints(ret)
	return ret
}
