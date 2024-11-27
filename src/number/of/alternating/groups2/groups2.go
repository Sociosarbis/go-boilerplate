package groups2

func numberOfAlternatingGroups(colors []int, k int) int {
	stack := make([]int, 0, k)
	stack = append(stack, 1)
	for i := 1; i < k; i++ {
		if colors[i] == colors[i-1] {
			stack[len(stack)-1]++
		} else {
			stack = append(stack, 1)
		}
	}
	var ret int
	if len(stack) == k {
		ret++
	}
	n := len(colors)
	ii := k - 1
	for i := 1; i < n; i++ {
		if stack[0] == 1 {
			stack = stack[1:]
		} else {
			stack[0]--
		}
		index := (i + k - 1) % n
		if colors[index] != colors[ii] {
			stack = append(stack, 1)
		} else {
			stack[len(stack)-1]++
		}
		if len(stack) == k {
			ret++
		}
		ii = index
	}
	return ret
}
