package operations12

func minOperations(nums []int) int {
	var out int
	stack := [][2]int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, [2]int{nums[i], 1})
		} else {
			top := stack[len(stack)-1]
			if top[0] == nums[i] {
				stack[len(stack)-1][1]++
			} else if top[0] < nums[i] {
				stack = append(stack, [2]int{nums[i], 1})
			} else {
				stack = stack[:len(stack)-1]
				out++
				i--
			}
		}
	}
	if stack[0][0] == 0 {
		return out + len(stack) - 1
	}
	return out + len(stack)
}
