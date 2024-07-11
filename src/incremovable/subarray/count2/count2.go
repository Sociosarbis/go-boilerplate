package count2

func incremovableSubarrayCount(nums []int) int64 {
	var ret int64
	n := len(nums)
	var left int
	rightStart := n - 1
	for ; rightStart > 0; rightStart-- {
		if nums[rightStart-1] >= nums[rightStart] {
			break
		}
	}
	for i := 0; i < n; i++ {
		ret++
		if rightStart < i+1 {
			rightStart = i + 1
		}
		for ; rightStart < n && nums[rightStart] <= left; rightStart++ {

		}
		ret += int64(n - rightStart)
		if nums[i] > left {
			left = nums[i]
		} else {
			break
		}
	}
	return ret
}
