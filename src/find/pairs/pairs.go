package pairs

func findPairs(nums []int, k int) int {
	m := map[int]int{}
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = 0
		}
		m[num]++
	}
	ret := 0
	for num := range m {
		if count, ok := m[num+k]; ok {
			if k == 0 {
				if count > 1 {
					ret += 1
				}
			} else {
				ret += 1
			}
		}
	}
	return ret
}
