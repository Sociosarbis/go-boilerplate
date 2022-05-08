package duplicates

func findDuplicates(nums []int) []int {
	ret := []int{}
	for i, num := range nums {
		for i != num-1 && num != 0 {
			if nums[num-1] == num {
				ret = append(ret, num)
				nums[i] = 0
				break
			} else {
				nums[num-1], nums[i] = nums[i], nums[num-1]
				num = nums[i]
			}
		}
	}
	return ret
}
