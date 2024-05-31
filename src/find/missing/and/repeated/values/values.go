package values

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	n = n*n + 1
	mp := make([]bool, n)
	ret := make([]int, 2)
	for _, row := range grid {
		for _, v := range row {
			if mp[v] {
				ret[0] = v
			} else {
				mp[v] = true
			}
		}
	}
	for i := 1; i < n; i++ {
		if !mp[i] {
			ret[1] = i
			return ret
		}
	}
	return ret
}
