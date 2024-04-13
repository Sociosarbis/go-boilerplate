package champion

func findChampion(n int, edges [][]int) int {
	losers := make([]bool, n)
	for _, edge := range edges {
		i := edge[1]
		if !losers[i] {
			losers[i] = true
		}
	}
	ret := -1
	for i := 0; i < n; i++ {
		if !losers[i] {
			if ret == -1 {
				ret = i
			} else {
				return -1
			}
		}
	}
	return ret
}
