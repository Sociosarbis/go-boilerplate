package lucky

func sumDigit(num int) int {
	ret := 0
	for num != 0 {
		ret += num % 10
		num /= 10
	}
	return ret
}

func getLucky(s string, k int) int {
	ret := 0
	for _, c := range s {
		ret += sumDigit(int(c) - 96)
	}
	for i := 2; i <= k; i++ {
		ret = sumDigit(ret)
	}
	return ret
}
