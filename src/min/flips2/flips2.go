package flips2

func getCount(grid [][]int, m, n, i, j int) (c int, a int, f int) {
	var b int
	ii := m - 1 - i
	jj := n - 1 - j
	if grid[i][j] == 1 {
		a++
	} else {
		b++
	}
	if ii != i {
		if grid[ii][j] == 1 {
			a++
		} else {
			b++
		}
		if jj != j {
			if grid[i][jj] == 1 {
				a++
			} else {
				b++
			}
			if grid[ii][jj] == 1 {
				a++
			} else {
				b++
			}
		}
	} else if jj != j {
		if grid[i][jj] == 1 {
			a++
		} else {
			b++
		}
	}
	c = a + b
	if a > b {
		a = c
		f = b
	} else {
		f = a
		a = 0
	}
	return c, a, f
}

func minFlips(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	e1 := m / 2
	e2 := n / 2
	if m%2 == 1 {
		e1++
	}
	if n%2 == 1 {
		e2++
	}
	var ret int
	var count int
	var hasFairFlip bool
	for i := 0; i < e1; i++ {
		for j := 0; j < e2; j++ {
			c, a, f := getCount(grid, m, n, i, j)
			ret += f
			count += a
			if !hasFairFlip && c == 2 && f == 1 {
				hasFairFlip = true
			}
		}
	}
	remain := count % 4
	switch remain {
	case 0:
		return ret
	case 1:
		return ret + 1
	case 2:
		if hasFairFlip {
			return ret
		}
		return ret + 2
	default:
		if hasFairFlip {
			return ret + 1
		}
		return ret + 3
	}
}
