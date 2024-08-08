package integer

func addedInteger(nums1 []int, nums2 []int) int {
	var max1 int
	var max2 int

	for _, num := range nums1 {
		if num > max1 {
			max1 = num
		}
	}

	for _, num := range nums2 {
		if num > max2 {
			max2 = num
		}
	}

	return max2 - max1
}
