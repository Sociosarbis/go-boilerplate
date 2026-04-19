package distance4

func maxDistance(nums1 []int, nums2 []int) int {
	var i, j int
	m, n := len(nums1), len(nums2)
	var out int
	for i < m && j < n {
		for i < m && nums1[i] > nums2[j] {
			i++
		}
		if i == m {
			break
		}
		if j >= i && j-i > out {
			out = j - i
		}
		j++
	}
	return out
}
