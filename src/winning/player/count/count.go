package count

func winningPlayerCount(n int, pick [][]int) int {
	counter := make([][11]int, n)

	for _, p := range pick {
		x, y := p[0], p[1]
		counter[x][y]++
	}
	var ret int
	for i := 0; i < n; i++ {
		for j := 0; j < 11; j++ {
			if counter[i][j] > i {
				ret++
				break
			}
		}
	}
	return ret
}
