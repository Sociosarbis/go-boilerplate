package sort

import "sort"

func wiggleSort(nums []int) {
	numsSort := make([]int, len(nums))
	copy(numsSort, nums)
	sort.Ints(numsSort)
	k := (len(nums) + 1) >> 1
	j := len(nums)
	for i := 0; i < len(nums); i++ {
		if i&1 == 0 {
			nums[i] = numsSort[k-1]
			k--
		} else {
			nums[i] = numsSort[j-1]
			j--
		}
	}
}
