package win

func canAliceWin(nums []int) bool {
	var ret int
	for _, num := range nums {
		if num > 9 {
			ret += num
		} else {
			ret -= num
		}
	}
	return ret != 0
}
