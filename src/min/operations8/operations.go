package operations8

func minOperations(nums []int, k int) int {
	var ret int
	for _, num := range nums {
		if num < k {
			ret++
		}
	}
	return ret
}
