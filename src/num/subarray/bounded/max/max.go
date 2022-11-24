package max

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	lp := make([]int, len(nums))
	lp[0] = -1
	for i := 1; i < len(nums); i++ {
		index := i - 1
		for index >= 0 && nums[index] < nums[i] {
			index = lp[index]
		}
		lp[i] = index
	}
	rp := make([]int, len(nums))
	rp[len(nums)-1] = len(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		index := i + 1
		for index < len(nums) && nums[index] <= nums[i] {
			index = rp[index]
		}
		rp[i] = index
	}
	ret := 0
	for i, num := range nums {
		if num >= left && num <= right {
			ret += (i - lp[i]) * (rp[i] - i)
		}
	}
	return ret
}
