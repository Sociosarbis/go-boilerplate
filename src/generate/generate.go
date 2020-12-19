package generate

// Pascal's Triangle
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	ret := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		for j := range row {
			row[j] = 0
		}
		prevRow := ret[i-1]
		for k := range ret[i-1] {
			row[k] += prevRow[k]
			row[k+1] += prevRow[k]
		}
		ret = append(ret, row)
	}
	return ret
}
