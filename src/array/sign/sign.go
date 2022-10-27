package sign

func arraySign(nums []int) int {
	ret := 1
	for _, num := range nums {
		if num < 0 {
			ret = -ret
		} else if num == 0 {
			return 0
		}
	}
	return ret
}
