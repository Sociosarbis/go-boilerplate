package array

func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	l := 1
	r := n
	var total int64 = int64(1+n) * int64(n) / 2
	for l <= r {
		mid := (l + r) / 2
		var left int64
		i := 0
		m := map[int]int{}
		j := 0
		for ; j < n; j++ {
			if c, ok := m[nums[j]]; ok {
				m[nums[j]] = c + 1
			} else {
				if len(m)+1 > mid {
					for i < j && len(m) == mid {
						if c, ok := m[nums[i]]; ok {
							if c == 1 {
								delete(m, nums[i])
							} else {
								m[nums[i]] = c - 1
							}
						}
						i++
					}
				}
				m[nums[j]] = 1
			}
			left += int64(j - i + 1)
		}
		if left < total-left {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
