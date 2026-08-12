package length

func maxSubarrayLength(nums []int, k int) int {
	n := len(nums)
	counter := make(map[int]int, n)
	var out int
	var l int
	for i, num := range nums {
		var count int
		if c, ok := counter[num]; ok {
			count = c + 1
		} else {
			count = 1
		}
		if i-l > out {
			out = i - l
		}
		if count > k {
			for ; l < i && count > k; l++ {
				if c, ok := counter[nums[l]]; ok {
					counter[nums[l]] = c - 1
					if nums[l] == num {
						count--
					}
				}
			}
		}
		counter[num] = count
	}
	if n-l > out {
		return n - l
	}
	return out
}
