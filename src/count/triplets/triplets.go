package triplets

const mask int64 = 1e9 + 7

func specialTriplets(nums []int) int {
	n := len(nums)
	prefix := make([]int, n)
	counter := make(map[int]int, n)
	for i, num := range nums {
		if c, ok := counter[num*2]; ok {
			prefix[i] = c
		}
		if c, ok := counter[num]; ok {
			counter[num] = c + 1
		} else {
			counter[num] = 1
		}
	}
	counter = make(map[int]int, n)
	var out int64
	for i := n - 1; i >= 1; i-- {
		if c, ok := counter[nums[i]*2]; ok {
			out = (out + int64(c)*int64(prefix[i])) % mask
		}
		if c, ok := counter[nums[i]]; ok {
			counter[nums[i]] = c + 1
		} else {
			counter[nums[i]] = 1
		}
	}
	return int(out)
}
