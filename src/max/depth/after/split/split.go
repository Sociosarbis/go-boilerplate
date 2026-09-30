package split

func maxDepthAfterSplit(seq string) []int {
	n := len(seq)
	stack := make([]byte, 0, n)
	out := make([]int, n)
	stack = append(stack, 0)
	for i := 1; i < n; i++ {
		if seq[i] == ')' {
			out[i] = int(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		} else {
			if len(stack) != 0 && stack[len(stack)-1] == 0 {
				out[i] = 1
				stack = append(stack, 1)
			} else {
				out[i] = 0
				stack = append(stack, 0)
			}
		}
	}
	return out
}
