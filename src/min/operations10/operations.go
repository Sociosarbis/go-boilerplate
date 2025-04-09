package operations10

func minOperations(nums []int, k int) int {
	exists := [101]bool{}
	var ret int
	for _, num := range nums {
		if num < k {
			return -1
		} else if num > k {
			if !exists[num] {
				exists[num] = true
				ret++
			}
		}
	}
	return ret
}
