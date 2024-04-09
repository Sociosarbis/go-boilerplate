package count

func maximumCount(nums []int) int {
	n := len(nums)
	l := 0
	r := n - 1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] < 0 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	neg := r + 1
	if r >= 0 {
		l = r
	} else {
		l = 0
	}
	r = n - 1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] <= 0 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	pos := n - l
	if pos > neg {
		return pos
	}
	return neg
}
