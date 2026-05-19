package common

func getCommon(nums1 []int, nums2 []int) int {
	var i, j int
	m, n := len(nums1), len(nums2)
	for i < m && j < n {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return -1
}
