package boats

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i := 0
	j := len(people) - 1
	var ret int
	for i < j {
		if people[j]+people[i] <= limit {
			i++
		}
		ret++
		j--
	}
	if i == j {
		ret++
	}
	return ret
}
