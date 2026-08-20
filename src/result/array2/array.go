package array2

func resultArray(nums []int) []int {
	n := len(nums)
	nums2 := make([]int, 0, n)
	a, b := nums[0], nums[1]
	nums2 = append(nums2, b)
	j := 1
	for i := 2; i < n; i++ {
		if a > b {
			a = nums[i]
			nums[j] = a
			j++
		} else {
			b = nums[i]
			nums2 = append(nums2, b)
		}
	}
	copy(nums[j:], nums2)
	return nums
}
