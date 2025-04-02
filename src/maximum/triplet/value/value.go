package value

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = max(suffix[i+1], nums[i])
	}
	e1 := n - 2
	e2 := n - 1
	var ret int64
	for i := 0; i < e1; i++ {
		a := nums[i]
		for j := i + 1; j < e2; j++ {
			if a < nums[j] {
				continue
			}
			b := int64(a-nums[j]) * int64(suffix[j+1])
			if b > ret {
				ret = b
			}
		}
	}
	return ret
}
