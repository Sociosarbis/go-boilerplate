package earnings

import (
	"sort"
)

type state struct {
	p int
	v int64
}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	sort.Slice(rides, func(i, j int) bool {
		if rides[i][0] < rides[j][0] {
			return true
		} else if rides[i][0] == rides[j][0] {
			return rides[i][1] < rides[j][1]
		}
		return false
	})

	stack := []state{{0, 0}}
	for _, ride := range rides {
		j := 0
		for j+1 < len(stack) && stack[j+1].p <= ride[0] {
			j++
		}
		if j != 0 {
			stack = stack[j:]
		}
		j = len(stack) - 1
		for stack[j].p > ride[0] {
			j--
		}
		item := state{ride[1], stack[j].v + int64(ride[1]-ride[0]) + int64(ride[2])}
		j = len(stack)
		for j > 0 && stack[j-1].p > item.p {
			if stack[j-1].v <= item.v {
				copy(stack[j-1:], stack[j:])
				stack = stack[:len(stack)-1]
			}
			j--
		}
		if j > 0 {
			if stack[j-1].v >= item.v {
				continue
			}
		}
		stack = append(stack, state{})
		copy(stack[j+1:], stack[j:])
		stack[j] = item
	}
	return stack[len(stack)-1].v
}
