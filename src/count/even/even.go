package even

func digitSum(num int) int {
	ret := 0
	for num != 0 {
		ret += num % 10
		num /= 10
	}
	return ret
}

func countEven(num int) int {
	ret := 0
	for i := 2; i <= num; i++ {
		sum := digitSum(i)
		if sum&1 == 0 && sum <= num {
			ret++
		}
	}
	return ret
}
