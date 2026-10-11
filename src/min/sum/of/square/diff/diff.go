package diff

import "sort"

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	k := int64(k1) + int64(k2)
	for i, a := range nums1 {
		nums1[i] = abs(a - nums2[i])
	}
	var out int64
	if k == 0 {
		for _, num := range nums1 {
			out += int64(num) * int64(num)
		}
		return out
	}
	sort.Ints(nums1)
	n := len(nums1)
	left := n - 1
	value := nums1[left]
	for left > 0 && nums1[left-1] == value {
		left--
	}
	for k > 0 && value != 0 {
		var prev int
		if left > 0 {
			prev = nums1[left-1]
		}
		diff := value - prev
		count := int64(n - left)
		needK := int64(diff) * int64(n-left)
		if needK <= k {
			k -= needK
			value = prev
			for left > 0 && nums1[left-1] == value {
				left--
			}
			if k == 0 || value == 0 {
				count = int64(n - left)
				out += count * int64(value) * int64(value)
			}
		} else {
			canSubtract := k / count
			value -= int(canSubtract)
			remain := k - canSubtract*count
			out += remain*int64(value-1)*int64(value-1) + int64(value)*int64(value)*(count-remain)
			break
		}
	}
	for i := 0; i < left; i++ {
		out += int64(nums1[i]) * int64(nums1[i])
	}
	return out
}
