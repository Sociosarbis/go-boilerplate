package sum

func maxAscendingSum(nums []int) int {
	ret := nums[0]
	temp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			temp += nums[i]
		} else {
			temp = nums[i]
		}
		if temp > ret {
			ret = temp
		}
	}
	return ret
}
