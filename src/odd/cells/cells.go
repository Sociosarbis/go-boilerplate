package cells

func oddCells(m int, n int, indices [][]int) int {
	rows := make([]int, m)
	cols := make([]int, n)
	for _, index := range indices {
		rows[index[0]]++
		cols[index[1]]++
	}

	oddCols := 0

	for _, col := range cols {
		if col&1 != 0 {
			oddCols++
		}
	}

	ret := 0

	for _, row := range rows {
		if row&1 != 0 {
			ret += n - oddCols
		} else {
			ret += oddCols
		}
	}

	return ret
}
