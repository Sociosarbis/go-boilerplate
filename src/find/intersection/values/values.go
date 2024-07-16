package values

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	group1, group2 := make([]bool, 101), make([]bool, 1001)
	for _, num := range nums1 {
		group1[num] = true
	}
	for _, num := range nums2 {
		group2[num] = true
	}
	ret := make([]int, 2)
	for _, num := range nums1 {
		if group2[num] {
			ret[0]++
		}
	}
	for _, num := range nums2 {
		if group1[num] {
			ret[1]++
		}
	}
	return ret
}
