package operations3

func minimumOperations(nums []int) int {
	var out int
	for _, num := range nums {
		if num%3 != 0 {
			out++
		}
	}
	return out
}
