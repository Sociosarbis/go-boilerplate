package cells

import "sort"

func maxIncreasingCells(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	items := make([][3]int, m*n)
	var index int
	for i, row := range mat {
		for j, v := range row {
			items[index] = [3]int{v, i, j}
			index++
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	})
	rowMax := make([][][2]int, m)
	colMax := make([][][2]int, n)

	for _, item := range items {
		v, r, c := item[0], item[1], item[2]
		t1 := 1
		rn := len(rowMax[r])
		for i := rn - 1; i >= 0; i-- {
			if rowMax[r][i][0] > v {
				t1 = rowMax[r][i][1] + 1
				break
			}
		}
		t2 := 1
		cn := len(colMax[c])
		for i := cn - 1; i >= 0; i-- {
			if colMax[c][i][0] > v {
				t2 = colMax[c][i][1] + 1
				break
			}
		}
		var max int
		if t1 > t2 {
			max = t1
		} else {
			max = t2
		}
		if rn == 0 || rowMax[r][rn-1][0] > v {
			rowMax[r] = append(rowMax[r], [2]int{v, max})
			if rn == 2 {
				rowMax[r] = rowMax[r][1:]
			}
		} else if rowMax[r][rn-1][1] < max {
			rowMax[r][rn-1][1] = max
		}
		if cn == 0 || colMax[c][cn-1][0] > v {
			colMax[c] = append(colMax[c], [2]int{v, max})
			if cn == 2 {
				colMax[c] = colMax[c][1:]
			}
		} else if colMax[c][cn-1][1] < max {
			colMax[c][cn-1][1] = max
		}
	}
	var ret int
	for _, row := range rowMax {
		for _, item := range row {
			if item[1] > ret {
				ret = item[1]
			}
		}
	}
	for _, col := range colMax {
		for _, item := range col {
			if item[1] > ret {
				ret = item[1]
			}
		}
	}
	return ret
}
