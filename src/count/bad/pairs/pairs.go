package pairs

func countBadPairs(nums []int) int64 {
	var ret int64
	n := len(nums)
	counter := make(map[int]int, n)
	counter[0] = 0
	for i := 1; i < n; i++ {
		d := (nums[i] - nums[0]) - i
		if c, ok := counter[d]; ok {
			counter[d] = c + 1
		} else {
			counter[d] = 1
		}
	}
	ret += int64(n - 1 - counter[0])
	for i := 1; i < n; i++ {
		d := (nums[i] - nums[0]) - i
		if c, ok := counter[d]; ok {
			counter[d] = c - 1
			ret += int64(n - i - c)
		}
	}
	return ret
}
