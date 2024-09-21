package score

func edgeScore(edges []int) int {
	n := len(edges)
	board := make([]int, n)

	for i, pointTo := range edges {
		board[pointTo] += i
	}

	ret := n - 1
	for i := n - 2; i >= 0; i-- {
		if board[i] >= board[ret] {
			ret = i
		}
	}
	return ret
}
