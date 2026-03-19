package submatrices

func numberOfSubmatrices(grid [][]byte) int {
	n := len(grid[0])
	prefix := make([][2][2]int, n)
	top := make([][2]int, n)
	var out int
	for i, row := range grid {
		index := i & 1
		left := [2]int{}
		if row[0] == 'X' {
			top[0][0]++
			left[0]++
		} else if row[0] == 'Y' {
			top[0][1]++
			left[1]++
		}
		if top[0][0] != 0 && top[0][0] == top[0][1] {
			out++
		}
		prefix[0][index][0], prefix[0][index][1] = top[0][0], top[0][1]
		for j := 1; j < n; j++ {
			if row[j] == 'X' {
				top[j][0]++
			} else if row[j] == 'Y' {
				top[j][1]++
			}
			x, y := top[j][0]+prefix[j-1][1-index][0]+left[0], top[j][1]+prefix[j-1][1-index][1]+left[1]
			prefix[j][index][0], prefix[j][index][1] = x, y
			if x != 0 && x == y {
				out++
			}
			if row[j] == 'X' {
				left[0]++
			} else if row[j] == 'Y' {
				left[1]++
			}
		}
	}
	return out
}
