package value4

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxValue(nums []int) []int {
	n := len(nums)
	prefixMax := make([]int, n)
	var temp int
	for i, num := range nums {
		temp = max(temp, num)
		prefixMax[i] = temp
	}
	out := make([]int, n)
	stack := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		l := len(stack)
		index := sort.Search(l, func(j int) bool {
			num := nums[stack[j]]
			return prefixMax[i] > num
		})
		if index < l {
			out[i] = max(prefixMax[i], out[stack[index]])
		} else {
			out[i] = prefixMax[i]
		}
		if l == 0 {
			stack = append(stack, i)
		} else if nums[stack[l-1]] > nums[i] {
			stack = append(stack, i)
		}
	}
	return out
}
