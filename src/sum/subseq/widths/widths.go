package widths

import "sort"

func sumSubseqWidths(nums []int) int {
	var ret int64
	n := len(nums)
	sort.Ints(nums)
	var mask int64 = 1e9 + 7
	var c int64 = 1
	for i, num := range nums {
		// 这里的最大最小值不是一一匹配的，只是它们会用到相同的c，所以组合到一起
		// 以num为最大值，以nums[n-i-1]为最小值
		ret = (ret + int64(num)*c - int64(nums[n-i-1])*c) % mask
		c = (c << 1) % mask
	}
	return int(ret)
}
