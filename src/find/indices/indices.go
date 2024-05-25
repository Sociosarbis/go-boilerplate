package indices

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	n := len(nums)
	for i, a := range nums {
		for j := i + indexDifference; j < n; j++ {
			b := nums[j]
			diff := a - b
			if diff >= valueDifference || diff <= -valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
