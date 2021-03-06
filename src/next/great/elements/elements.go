package elements

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	extendedNums := append(nums, nums[:len(nums)-1]...)

	stack := []int{}
	res := make([]int, len(nums))

	for i := 0; i < len(extendedNums); i++ {
		// 当遇到比自己大的数时出栈
		for len(stack) != 0 && extendedNums[i] > nums[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = extendedNums[i]
			stack = stack[:len(stack)-1]
		}
		if i < len(nums) {
			stack = append(stack, i)
		}
	}

	for i := 0; i < len(stack); i++ {
		res[stack[i]] = -1
	}
	return res
}
