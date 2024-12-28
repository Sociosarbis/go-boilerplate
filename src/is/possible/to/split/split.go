package split

func isPossibleToSplit(nums []int) bool {
	counter := make([]byte, 101)
	for _, num := range nums {
		if counter[num] >= 2 {
			return false
		}
		counter[num]++
	}
	return true
}
