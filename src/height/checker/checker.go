package checker

import "sort"

func heightChecker(heights []int) int {
	sortedHeights := make([]int, len(heights))
	copy(sortedHeights, heights)
	sort.Ints(sortedHeights)
	ret := 0
	for i, height := range heights {
		if sortedHeights[i] != height {
			ret++
		}
	}
	return ret
}
