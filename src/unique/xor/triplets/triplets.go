package triplets

func uniqueXorTriplets(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	var count int
	for n != 0 {
		n >>= 1
		count++
	}
	return 1 << count
}
