package boxes2

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	sort.Ints(capacity)
	var sum int
	for _, it := range apple {
		sum += it
	}
	i := len(capacity) - 1
	for sum > 0 && i >= 0 {
		sum -= capacity[i]
		i--
	}
	return len(capacity) - i - 1
}
