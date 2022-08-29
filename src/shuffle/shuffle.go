package shuffle

func shuffle(nums []int, n int) []int {
	ret := make([]int, n*2)
	for i := 0; i < len(ret); i++ {
		if i < n {
			ret[i*2] = nums[i]
		} else {
			ret[(i-n)*2+1] = nums[i]
		}
	}
	return ret
}
