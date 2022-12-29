package three

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m := make([]byte, 101)
	for _, num := range nums1 {
		m[num] = 1
	}
	for _, num := range nums2 {
		m[num] |= 2
	}
	for _, num := range nums3 {
		m[num] |= 4
	}
	ret := []int{}
	for n, c := range m {
		if c != 0 && c&(c-1) != 0 {
			ret = append(ret, n)
		}
	}
	return ret
}
