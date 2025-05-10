package sum

func count(nums []int) (int64, int) {
	var sum int64
	var count int
	for _, num := range nums {
		if num == 0 {
			count++
		} else {
			sum += int64(num)
		}
	}
	return sum, count
}

func minSum(nums1 []int, nums2 []int) int64 {
	s1, c1 := count(nums1)
	s2, c2 := count(nums2)
	m1, m2 := s1+int64(c1), s2+int64(c2)
	if m1 == m2 {
		return m1
	}
	if m1 > m2 {
		if c2 == 0 {
			return -1
		}
		return m1
	}
	if c1 == 0 {
		return -1
	}
	return m2
}
