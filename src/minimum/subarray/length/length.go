package length

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	ret := n + 1
	for i := 0; i < n; i++ {
		var temp int
		for j := i; j < n; j++ {
			temp |= nums[j]
			if temp >= k {
				if j-i+1 < ret {
					ret = j - i + 1
				}
				break
			}
		}
	}
	if ret == n+1 {
		return -1
	}
	return ret
}
