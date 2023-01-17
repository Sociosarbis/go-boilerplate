package pairs

func diff(num int) int64 {
	a := int64(num)
	var b int64
	var base int64 = 1
	digits := []int{}
	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		b += base * int64(digits[i])
		base *= 10
	}
	return a - b
}

func countNicePairs(nums []int) int {
	m := map[int64]int{}
	for _, num := range nums {
		d := diff(num)
		if c, ok := m[d]; ok {
			m[d] = c + 1
		} else {
			m[d] = 1
		}
	}

	ret := 0
	var mask int = 1e9 + 7
	for _, c := range m {
		if c > 1 {
			ret = (ret + int((int64(c)*(int64(c)-1)/2)%int64(mask))) % mask
		}
	}
	return ret
}
