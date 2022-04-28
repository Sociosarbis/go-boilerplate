package parity

func sortArrayByParity(nums []int) []int {
	i := len(nums) - 1
	j := 0
	for j < i {
		for i > j && nums[i]&1 != 0 {
			i--
		}
		for j < i && nums[j]&1 == 0 {
			j++
		}
		nums[j], nums[i] = nums[i], nums[j]

	}
	return nums
}
