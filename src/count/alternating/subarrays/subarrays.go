package subarrays

func countAlternatingSubarrays(nums []int) int64 {
	n := len(nums)
	ret := int64(n)

	var start int
	i := 1
	for ; i < n; i++ {
		if nums[i] == nums[i-1] {
			if i-start > 1 {
				ret += int64(i-start) * int64(i-start-1) / 2
			}
			start = i
		}
	}

	if i-start > 1 {
		ret += int64(i-start) * int64(i-start-1) / 2
	}

	return ret
}
