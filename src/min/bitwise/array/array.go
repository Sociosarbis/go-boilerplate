package array

func minBitwiseArray(nums []int) []int {
	for i, num := range nums {
		if num == 2 {
			nums[i] = -1
		} else {
			base := 1
			for num&(base<<1) != 0 {
				base <<= 1
			}
			nums[i] = num - base
		}
	}
	return nums
}
