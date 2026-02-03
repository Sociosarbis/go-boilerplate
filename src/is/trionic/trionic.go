package trionic

func isTrionic(nums []int) bool {
	n := len(nums)
	if nums[1] <= nums[0] || nums[n-1] <= nums[n-2] {
		return false
	}
	var count int
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			if count == 1 {
				count++
			}
		} else if nums[i] < nums[i-1] {
			if count != 1 {
				if count == 2 {
					return false
				}
				count++
			}
		} else {
			return false
		}
	}
	return count == 2
}
