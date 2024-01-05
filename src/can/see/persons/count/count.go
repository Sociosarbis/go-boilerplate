package count

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	ret := make([]int, n)
	stack := make([]int, 0, n)
	stack = append(stack, heights[n-1])

	for i := n - 2; i >= 0; i-- {
		r := len(stack) - 1
		for ; r > 0; r-- {
			if stack[r] > heights[i] {
				break
			}
		}
		ret[i] = len(stack) - r
		if stack[r] > heights[i] {
			stack = stack[:r+1]
		} else {
			stack = stack[:r]
		}
		stack = append(stack, heights[i])
	}
	return ret
}
