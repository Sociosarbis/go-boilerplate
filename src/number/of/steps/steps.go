package steps

func numberOfSteps(num int) int {
	ret := 0
	for num != 0 {
		if num&1 == 1 {
			num -= 1
		} else {
			num >>= 1
		}
		ret += 1
	}
	return ret
}
