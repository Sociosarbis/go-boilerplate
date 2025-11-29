package operations14

func minOperations(nums []int, k int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum % k
}
