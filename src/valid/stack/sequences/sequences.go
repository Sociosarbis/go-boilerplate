package sequences

func validateStackSequences(pushed []int, popped []int) bool {
	i := 0
	j := 0
	stack := []int{}
	for j < len(popped) {
		if len(stack) == 0 || stack[len(stack)-1] != popped[j] {
			if i < len(pushed) {
				stack = append(stack, pushed[i])
				i++
			} else {
				return false
			}
		} else {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return true
}
