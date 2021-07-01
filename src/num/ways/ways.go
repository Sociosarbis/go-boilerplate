package ways

func numWays(n int, relation [][]int, k int) int {
	numToNext := make([][]int, n)
	for _, r := range relation {
		numToNext[r[0]] = append(numToNext[r[0]], r[1])
	}
	queue := [2][]int{}
	for i := 0; i < 2; i += 1 {
		queue[i] = make([]int, n)
		for j := 0; j < n; j += 1 {
			queue[i][j] = 0
		}
	}
	queue[0][0] = 1
	for step := 0; step <= k; step += 1 {
		index := step & 1
		for i := 0; i < n; i += 1 {
			if queue[index][i] != 0 {
				for _, next := range numToNext[i] {
					queue[1-index][next] += queue[index][i]
				}
			}
		}
		if step != k {
			for i := 0; i < n; i += 1 {
				queue[index][i] = 0
			}
		}
	}
	return queue[k&1][n-1]
}
