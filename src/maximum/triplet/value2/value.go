package value2

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	suffix := make([]int, n)
	for i := n - 2; i > 0; i-- {
		suffix[i] = max(suffix[i+1], nums[i+1])
	}
	e := n - 1
	var prefix int
	var ret int64
	for i := 1; i < e; i++ {
		prefix = max(prefix, nums[i-1])
		if prefix > nums[i] {
			v := int64(prefix-nums[i]) * int64(suffix[i])
			if v > ret {
				ret = v
			}
		}
	}
	return ret
}
