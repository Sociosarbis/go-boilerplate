package count

import "sort"

type tuple struct {
	i   int
	num int
}

type tuples []tuple

func (this tuples) Len() int {
	return len(this)
}

func (this tuples) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this tuples) Less(i, j int) bool {
	if this[i].num < this[j].num {
		return true
	} else if this[i].num == this[j].num {
		return this[i].i < this[j].i
	}
	return false
}

func advantageCount(nums1 []int, nums2 []int) []int {
	nums1Copy := make([]int, len(nums1))
	copy(nums1Copy, nums1)
	sort.Ints(nums1Copy)
	arr := make([]tuple, len(nums2))
	for i, num := range nums2 {
		arr[i].i = i
		arr[i].num = num
	}
	ret := make([]int, len(nums1))
	sort.Sort(tuples(arr))
	l := 0
	r := len(arr) - 1
	for i := r; i >= 0; i-- {
		if nums1Copy[r] > arr[i].num {
			ret[arr[i].i] = nums1Copy[r]
			r--
		} else {
			ret[arr[i].i] = nums1Copy[l]
			l++
		}
	}
	return ret
}
