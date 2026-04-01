package healths

import (
	"sort"
)

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return positions[indices[i]] < positions[indices[j]]
	})
	stack := []int{}
	for i := 0; i < n; i++ {
		m := len(stack)
		if m == 0 {
			stack = append(stack, i)
		} else {
			left, right := indices[stack[m-1]], indices[i]
			if directions[left] == 'R' {
				if directions[right] == 'L' {
					if healths[left] < healths[right] {
						stack = stack[:m-1]
						healths[right]--
						if healths[right] > 0 {
							i--
							continue
						}
					} else if healths[left] > healths[right] {
						healths[left]--
						if healths[left] == 0 {
							stack = stack[:m-1]
						}
					} else {
						stack = stack[:m-1]
					}
				} else {
					stack = append(stack, i)
				}
			} else {
				stack = append(stack, i)
			}
		}
	}
	for i, index := range stack {
		stack[i] = indices[index]
	}
	sort.Ints(stack)
	for i, index := range stack {
		stack[i] = healths[index]
	}
	return stack
}
