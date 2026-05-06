package box

func rotateTheBox(boxGrid [][]byte) [][]byte {
	m, n := len(boxGrid), len(boxGrid[0])
	end := len(boxGrid[0]) - 1
	for _, row := range boxGrid {
		k := end
		for j := end; j >= 0; j-- {
			if row[j] == '#' {
				row[j], row[k] = row[k], row[j]
				k--
			} else if row[j] == '*' {
				k = j - 1
			}
			if row[j] != '.' {
				for ; k >= 0; k-- {
					if row[k] == '.' {
						break
					}
				}
				j = k
			}
		}
	}
	out := make([][]byte, n)
	for i := 0; i < n; i++ {
		out[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			out[i][j] = boxGrid[m-j-1][i]
		}
	}
	return out
}
