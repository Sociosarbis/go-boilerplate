package subarray

func getS(num int) byte {
	if num < 0 {
		return 2
	}
	if num > 0 {
		return 1
	}
	return 0
}

func longestMonotonicSubarray(nums []int) int {
	n := len(nums)
	var s byte
	var out, left int
	for i := 1; i < n; i++ {
		ns := getS(nums[i] - nums[i-1])
		if ns == 0 || (s != 0 && ns != s) {
			if i-left > out {
				out = i - left
			}
			if ns != 0 {
				left = i - 1
			} else {
				left = i
			}
		}
		s = ns
	}
	if n-left > out {
		out = n - left
	}
	return out
}
