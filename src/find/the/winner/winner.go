package winner

func findTheWinner(n int, k int) int {
	players := make([]int, n)

	var i int

	for i = 0; i < n; i++ {
		players[i] = i + 1
	}

	i = 0

	for len(players) > 1 {
		i = (i + k - 1) % len(players)
		if i+1 < len(players) {
			copy(players[i:], players[i+1:])
		}
		players = players[:len(players)-1]
	}
	return players[0]
}
