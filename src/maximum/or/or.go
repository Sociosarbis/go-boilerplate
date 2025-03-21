package or

func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] | nums[i]
	}
	var ret int64
	var temp int
	for i, num := range nums {
		v := (int64(num) << k) | int64(temp) | int64(suffix[i+1])
		if v > ret {
			ret = v
		}
		temp |= num
	}
	return ret
}
