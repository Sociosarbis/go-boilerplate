package marbles

import "sort"

type empty struct{}

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	positions := make(map[int]empty, len(nums))
	for _, num := range nums {
		if _, ok := positions[num]; !ok {
			positions[num] = empty{}
		}
	}

	for i, from := range moveFrom {
		if _, ok := positions[from]; ok {
			delete(positions, from)
			positions[moveTo[i]] = empty{}
		}
	}

	ret := make([]int, 0, len(positions))
	for i := range positions {
		ret = append(ret, i)
	}
	sort.Ints(ret)
	return ret
}
