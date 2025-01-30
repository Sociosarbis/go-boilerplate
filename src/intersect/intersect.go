package intersect

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var i int
	var n int
	n1 := len(nums1)
	n2 := len(nums2)
	if n1 < n2 {
		n = n1
	} else {
		n = n2
	}
	ret := make([]int, n)
	var j int
	for i < n1 && j < n2 {
		if nums1[i] == nums2[j] {
			ret = append(ret, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return ret
}
