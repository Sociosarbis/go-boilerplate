package subarrays

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
