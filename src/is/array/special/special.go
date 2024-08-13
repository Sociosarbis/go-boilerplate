package special

func isArraySpecial(nums []int) bool {
	isOdd := nums[0]%2 == 1
	n := len(nums)
	for i := 1; i < n; i++ {
		if isOdd != (nums[i]%2 == 1) {
			isOdd = !isOdd
		} else {
			return false
		}
	}
	return true
}
