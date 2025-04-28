package subarrays3

func countSubarrays(nums []int, k int64) int64 {
	var j int
	var temp int64
	n := len(nums)
	var ret int64
	for i := 0; i < n; i++ {
		if j < i {
			j = i
		}
		for j < n && (temp+int64(nums[j]))*int64(j-i+1) < k {
			temp += int64(nums[j])
			j++
		}
		if j > i {
			ret += int64(j - i)
			temp -= int64(nums[i])
		}
	}
	return ret
}
