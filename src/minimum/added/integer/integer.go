package integer

import "sort"

func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	n1 := len(nums1)
	n2 := len(nums2)

	if n2 == 1 {
		return nums2[0] - nums1[n1-1]
	}

loop:
	for i := 0; i < 2; i++ {
		prev := nums1[n1-i-1]
		cost := i
		ii := n2 - 2
		for j := n1 - i - 2; j >= 0; j-- {
			if prev-nums1[j] != nums2[ii+1]-nums2[ii] {
				cost++
				if cost > 2 {
					continue loop
				} else {
					continue
				}
			} else {
				prev = nums1[j]
			}
			ii--
			if ii < 0 {
				break
			}
		}
		return nums2[n2-1] - nums1[n1-i-1]
	}
	return nums2[n2-1] - nums1[n1-3]
}
