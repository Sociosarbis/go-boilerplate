package operations

func maxOperations(nums []int) int {
	ret := 1
	sum := nums[0] + nums[1]
	n := len(nums)

	for i := 3; i < n; i += 2 {
		if nums[i]+nums[i-1] == sum {
			ret++
		} else {
			break
		}
	}
	return ret
}
