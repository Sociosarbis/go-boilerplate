package utils

func Min(num ...int) int {
	min := num[0]
	for _, n := range num[1:] {
		if n < min {
			min = n
		}
	}
	return min
}

func Max(nums ...int) int {
	ret := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] > ret {
			ret = nums[i]
		}
	}
	return ret
}
