package numbers

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			tmp := nums[i]
			for tmp != i+1 && tmp != nums[tmp-1] {
				nextTmp := nums[tmp-1]
				nums[tmp-1] = tmp
				tmp = nextTmp
			}
			nums[i] = tmp
		}
	}
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ret = append(ret, i+1)
		}
	}
	return ret
}
