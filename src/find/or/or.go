package or

func findKOr(nums []int, k int) int {
	ret := 0
	for i := 0; i <= 30; i++ {
		temp := 0
		mask := 1 << i
		for _, num := range nums {
			if num&mask == mask {
				temp++
			}
		}
		if temp >= k {
			ret |= mask
		}
	}
	return ret
}
