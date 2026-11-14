package subarray

func minimumSumSubarray(nums []int, l int, r int) int {
	var prefix int
	for i := l - 2; i >= 0; i-- {
		prefix += nums[i]
	}
	n := len(nums)
	e1 := n - l
	out := -1
	for i := 0; i <= e1; i++ {
		if i > 0 {
			prefix -= nums[i-1]
		}
		var e2 int
		if i+r > n {
			e2 = n - i
		} else {
			e2 = r
		}
		temp := prefix
		for j := l - 1; j < e2; j++ {
			temp += nums[i+j]
			if temp > 0 && (out == -1 || temp < out) {
				out = temp
			}
		}
		prefix += nums[i+l-1]
	}
	return out
}
