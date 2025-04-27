package subarrays2

func countSubarrays(nums []int) int {
	end := len(nums) - 1
	var ret int
	for i := 1; i < end; i++ {
		if nums[i] == (nums[i-1]+nums[i+1])<<1 {
			ret++
		}
	}
	return ret
}
