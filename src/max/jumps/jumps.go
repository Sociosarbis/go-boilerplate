package jumps

func maxJumps(arr []int, d int) int {
	n := len(arr)
	count := make([]int, n)
	next := make([][]int, n)
	stack := make([]int, 0, n)
	stack = append(stack, 0)
	for i := 1; i < n; i++ {
		for len(stack) != 0 && i-stack[0] > d {
			stack = stack[1:]
		}
		for len(stack) != 0 && arr[i] >= arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		if len(stack) > 1 {
			nextIndex := stack[len(stack)-2]
			next[i] = append(next[i], nextIndex)
			count[nextIndex]++
		}
	}
	stack = make([]int, 0, n)
	stack = append(stack, n-1)
	for i := n - 2; i >= 0; i-- {
		for len(stack) != 0 && stack[0]-i > d {
			stack = stack[1:]
		}
		for len(stack) != 0 && arr[i] >= arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		if len(stack) > 1 {
			nextIndex := stack[len(stack)-2]
			next[i] = append(next[i], nextIndex)
			count[nextIndex]++
		}
	}
	var out int
	stack = make([]int, 0, n)
	for i := 0; i < n; i++ {
		if count[i] == 0 {
			stack = append(stack, i)
		}
	}
	l := len(stack)
	for l != 0 {
		out++
		for j := 0; j < l; j++ {
			index := stack[j]
			for _, i := range next[index] {
				count[i]--
				if count[i] == 0 {
					stack = append(stack, i)
				}
			}
		}
		stack = stack[l:]
		l = len(stack)
	}
	return out
}
