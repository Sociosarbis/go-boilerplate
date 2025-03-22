package ones

func rowAndMaximumOnes(mat [][]int) []int {
	ret := make([]int, 2)
	for _, num := range mat[0] {
		if num == 1 {
			ret[1]++
		}
	}
	n := len(mat)
	for i := 1; i < n; i++ {
		var temp int
		for _, num := range mat[i] {
			if num == 1 {
				temp++
			}
		}
		if temp > ret[1] {
			ret[0], ret[1] = i, temp
		}
	}
	return ret
}
