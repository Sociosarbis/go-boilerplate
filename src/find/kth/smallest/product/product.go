package product

import "sort"

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	n1, n2 := len(nums1), len(nums2)
	a, b, c, d := int64(nums1[0])*int64(nums2[n2-1]), int64(nums1[0])*int64(nums2[0]), int64(nums1[n1-1])*int64(nums2[0]), int64(nums1[n1-1])*int64(nums2[n2-1])
	l := min(min(min(a, b), c), d)
	r := max(max(max(a, b), c), d)
	out := int64(n1) * int64(n2)
loop:
	for l <= r {
		mid := (l + r) >> 1
		var count int64
		for i := 0; i < n1; i++ {
			a := int64(nums1[i])
			var c int
			if a < 0 {
				c = sort.Search(n2, func(i int) bool {
					return int64(nums2[i])*a <= mid
				})
				count += int64(n2 - c)
			} else {
				c = sort.Search(n2, func(i int) bool {
					return int64(nums2[i])*a > mid
				})
				count += int64(c)
			}
			if count >= k {
				out = mid
				r = mid - 1
				continue loop
			}
		}
		l = mid + 1
	}
	return out
}
