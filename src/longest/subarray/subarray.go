package subarray

func longestSubarray(nums []int) int {
	out := [2]int{nums[0], 1}
	count := 1
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] > out[0] {
			count = 1
			out[0], out[1] = nums[i], 1
		} else if nums[i] == out[0] {
			count++
		} else {
			if out[1] < count {
				out[1] = count
			}
			if count > 0 {
				count = 0
			}
		}
	}
	if out[1] < count {
		return count
	}
	return out[1]
}
