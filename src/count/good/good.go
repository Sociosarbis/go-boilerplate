package good

func countGood(nums []int, k int) int64 {
	counter := map[int]int{}
	n := len(nums)
	var j int
	var temp int
	var ret int64
	for i := 0; i < n; i++ {
		for temp < k && j < n {
			if c, ok := counter[nums[j]]; ok {
				temp += c
				counter[nums[j]] = c + 1
			} else {
				counter[nums[j]] = 1
			}
			j++
		}
		if temp < k {
			break
		}
		ret += int64(n - j + 1)
		if c, ok := counter[nums[i]]; ok && c > 0 {
			temp -= c - 1
			counter[nums[i]] = c - 1
		}
	}
	return ret
}
