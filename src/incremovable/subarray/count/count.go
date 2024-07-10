package count

func incremovableSubarrayCount(nums []int) int {
	var ret int
	n := len(nums)
	var left int
	var right int
	for i := 0; i < n; i++ {
		right = 0
		ret++
		for j := n - 1; j > i; j-- {
			if nums[j] > left && (right == 0 || nums[j] < right) {
				right = nums[j]
				ret++
			} else {
				break
			}
		}
		if nums[i] > left {
			left = nums[i]
		} else {
			break
		}
	}
	return ret
}
