package subarray2

func longestSubarray(nums []int) int {
	n := len(nums)
	var j int
	for i := 1; i < n; i++ {
		if nums[i] == 1 {
			if nums[j] == 0 {
				j++
				nums[j] = 1
			} else {
				nums[j]++
			}
		} else {
			j++
			nums[j] = 0
		}
	}
	if j == 0 {
		if nums[0] == 0 {
			return 0
		}
		return nums[0] - 1
	}
	var out int
	for i := 0; i <= j; i++ {
		if i > 1 && nums[i] > 0 && nums[i-2] > 0 && nums[i]+nums[i-2] > out {
			out = nums[i] + nums[i-2]
		} else if nums[i] > out {
			out = nums[i]
		}
	}
	return out
}
