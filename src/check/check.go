package check

func check(nums []int) bool {
	hasRot := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if hasRot {
				return false
			}
			hasRot = true
		}
	}
	if !hasRot {
		return true
	}
	return nums[0] >= nums[len(nums)-1]
}
