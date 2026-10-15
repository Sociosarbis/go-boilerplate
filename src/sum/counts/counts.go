package counts

func sumCounts(nums []int) int {
	var counter [101]int
	var count, out int
	n := len(nums)
	r := -1
	for i := 0; i < n; i++ {
		if i > 0 {
			counter[nums[i-1]]--
			if counter[nums[i-1]] == 0 {
				count--
			}
		}
		if r < i {
			r = i
			counter[nums[i]]++
			if counter[nums[i]] == 1 {
				count++
			}
		}
		out += count * count
		if r+1 == n {
			for ; r > i; r-- {
				counter[nums[r]]--
				if counter[nums[r]] == 0 {
					count--
				}
				out += count * count
			}
		} else {
			for ; r+1 < n; r++ {
				counter[nums[r+1]]++
				if counter[nums[r+1]] == 1 {
					count++
				}
				out += count * count
			}
		}
	}
	return out
}
