package operations6

func minOperations(nums []int) int {
	var i int
	n := len(nums)
	var ret int
	end := n - 2
	for i < end {
		if nums[i] != 1 {
			nums[i+1] = 1 - nums[i+1]
			nums[i+2] = 1 - nums[i+2]
			ret++
		}
	}
	if nums[n-1] != 1 || nums[n-2] != 1 {
		return -1
	}
	return ret
}
