package integer

func largestInteger(nums []int, k int) int {
	n := len(nums)
loop:
	for i := 50; i >= 0; i-- {
		var count int
		for j := n - k; j >= 0; j-- {
			for l := 0; l < k; l++ {
				if nums[j+l] == i {
					count++
					break
				}
			}
			if count > 1 {
				continue loop
			}
		}
		if count == 1 {
			return i
		}
	}
	return -1
}
