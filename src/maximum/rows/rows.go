package rows

func maximumRows(matrix [][]int, numSelect int) int {
	bitRows := make([]int, len(matrix))
	for i, row := range matrix {
		for j, num := range row {
			if num == 1 {
				bitRows[i] |= 1 << j
			}
		}
	}
	return dfs(bitRows, 0, 0, numSelect)
}

func dfs(bitRows []int, i, temp, numSelect int) int {
	if i >= len(bitRows) {
		return 0
	}
	var res1 int
	if countBits(bitRows[i]|temp) <= numSelect {
		res1 = 1 + dfs(bitRows, i+1, bitRows[i]|temp, numSelect)
	}
	res2 := dfs(bitRows, i+1, temp, numSelect)
	if res1 >= res2 {
		return res1
	}
	return res2
}

func countBits(bits int) int {
	ret := 0
	for bits != 0 {
		bits &= bits - 1
		ret++
	}
	return ret
}
