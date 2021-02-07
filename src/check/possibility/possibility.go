package possibility

func checkPossibility(nums []int) bool {
	hasModified := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if hasModified {
				return false
			}
			m := nums[i]
			hasModified = true
			if i-1 > 0 {
				if nums[i-2] > m {
					nums[i] = nums[i-1]
					continue
				}
			}
			nums[i-1] = m
		}
	}
	return true
}
