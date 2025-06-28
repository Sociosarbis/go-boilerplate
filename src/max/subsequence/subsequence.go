package subsequence

import "sort"

func maxSubsequence(nums []int, k int) []int {
	arr := make([][2]int, len(nums))
	for i, num := range nums {
		arr[i][0], arr[i][1] = i, num
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	arr = arr[:k]
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	for i := 0; i < k; i++ {
		nums[i] = arr[i][1]
	}
	return nums[:k]
}
