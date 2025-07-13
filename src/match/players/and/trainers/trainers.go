package trainers

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(trainers)
	sort.Ints(players)
	var out, j int
	n := len(trainers)
	for _, player := range players {
		for j < n && player > trainers[j] {
			j++
		}
		if j < n {
			out++
			j++
		} else {
			break
		}
	}
	return out
}
