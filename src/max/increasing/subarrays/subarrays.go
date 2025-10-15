package subarrays2

func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	var start int
	lastEnd := -1
	i := 1
	for i < n {
		if nums[i] > nums[i-1] {
			i++
			l := i - start
			if l >= 2*k || (l >= k && lastEnd != -1 && lastEnd+1 == start) {
				return true
			}
		} else {
			if i-start >= k {
				lastEnd = i - 1
			}
			start = i
			i++
		}
	}
	return i-start >= k && lastEnd != -1 && lastEnd+1 == start
}

func maxIncreasingSubarrays(nums []int) int {
	l, r := 1, len(nums)/2
	out := 1
	for l <= r {
		mid := (l + r) >> 1
		if hasIncreasingSubarrays(nums, mid) {
			l = mid + 1
			if mid > out {
				out = mid
			}
		} else {
			r = mid - 1
		}
	}
	return out
}
