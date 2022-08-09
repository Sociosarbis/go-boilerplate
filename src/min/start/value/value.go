package value

func minStartValue(nums []int) int {
	ret := 0
	temp := 0
	for _, num := range nums {
		temp += num
		if temp < ret {
			ret = temp
		}
	}
	if ret < 1 {
		return 1 - ret
	} else {
		return 1
	}
}
