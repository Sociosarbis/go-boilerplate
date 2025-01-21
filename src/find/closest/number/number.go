package number

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func findClosestNumber(nums []int) int {
	ret := nums[0]
	n := len(nums)
	var r int
	for i := 1; i < n; i++ {
		r = abs(ret) - abs(nums[i])
		if r > 0 {
			ret = nums[i]
		} else if r == 0 && nums[i] > ret {
			ret = nums[i]
		}
	}
	return ret
}
