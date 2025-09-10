package teachings

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)
	users := make([][]bool, m+1)
	for i, it := range languages {
		users[i+1] = make([]bool, n+1)
		for _, j := range it {
			users[i+1][j] = true
		}
	}
	var i int
loop:
	for _, friendship := range friendships {
		a, b := friendship[0], friendship[1]
		for _, it := range languages[a-1] {
			if users[b][it] {
				continue loop
			}
		}
		for _, it := range languages[b-1] {
			if users[a][it] {
				continue loop
			}
		}
		friendships[i] = friendship
		i++
	}
	friendships = friendships[:i]
	if len(friendships) == 0 {
		return 0
	}
	out := m
	for i := 1; i <= n; i++ {
		var restore []int
		for _, friendship := range friendships {
			a, b := friendship[0], friendship[1]
			if !users[a][i] {
				restore = append(restore, a)
				users[a][i] = true
			}
			if !users[b][i] {
				restore = append(restore, b)
				users[b][i] = true
			}
		}
		if len(restore) < out {
			out = len(restore)
		}
		for _, it := range restore {
			users[it][i] = false
		}
	}
	return out
}
