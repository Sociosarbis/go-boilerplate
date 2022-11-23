package balls

func getIndex(num int) int {
	ret := 0
	for num != 0 {
		ret += num % 10
		num /= 10
	}
	return ret
}

func countBalls(lowLimit int, highLimit int) int {
	counter := make([]int, 46)
	for i := lowLimit; i <= highLimit; i++ {
		counter[getIndex(i)]++
	}
	ret := 0
	for i := 0; i < 46; i++ {
		if counter[i] > ret {
			ret = counter[i]
		}
	}
	return ret
}
