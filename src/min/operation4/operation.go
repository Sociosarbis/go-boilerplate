package operation4

func minOperations(nums []int) int {
	temp := nums[0]
	ret := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= temp {
			temp++
			ret += temp - nums[i]
		} else {
			temp = nums[i]
		}
	}
	return ret
}
