package subarray

func beautifulSubarrays(nums []int) int64 {
	n := len(nums)
	counter := make(map[int]int, n)
	var mask int
	for _, num := range nums {
		mask ^= num
		if c, ok := counter[mask]; ok {
			counter[mask] = c + 1
		} else {
			counter[mask] = 1
		}
	}
	mask = 0
	var ret int64
	if c, ok := counter[mask]; ok {
		ret += int64(c)
	}
	for _, num := range nums {
		mask ^= num
		if c, ok := counter[mask]; ok {
			c--
			ret += int64(c)
			if c == 0 {
				delete(counter, mask)
			} else {
				counter[mask] = c
			}
		}
	}
	return ret
}
