package i

func smallestRangeI(nums []int, k int) int {
	min := 10001
	max := -1

	for _, num := range nums {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	if max-min > 2*k {
		return max - min - 2*k
	}
	return 0
}
