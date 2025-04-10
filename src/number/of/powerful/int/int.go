package int

func getDigits(start int64) [16]byte {
	var ret [16]byte
	for i := 0; start != 0; i++ {
		ret[i] = byte(start % 10)
		start /= 10
	}
	return ret
}

func pow(num int64, n int) int64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return num
	}
	return pow(num, n/2) * pow(num, n-n/2)
}

func count(upperDigits, digits [16]byte, from, limit, n int) int64 {
	for i := from; i >= 0; i-- {
		if upperDigits[i] == 0 && digits[i] == 0 {
			continue
		}
		if i < n {
			for ; i >= 0; i-- {
				if upperDigits[i] > digits[i] {
					return 1
				} else if upperDigits[i] < digits[i] {
					return 0
				}
			}
			return 1
		}
		if upperDigits[i] < digits[i] {
			return 0
		} else {
			max := upperDigits[i]
			if max > byte(limit) {
				max = byte(limit)
			}
			if max < digits[i] {
				return 0
			}
			if max == upperDigits[i] {
				return int64(max-digits[i])*pow(int64(limit+1), i-n) + count(upperDigits, digits, i-1, limit, n)
			} else {
				return int64(max-digits[i]+1) * pow(int64(limit+1), i-n)
			}
		}
	}
	return 0
}

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	upperDigits1 := getDigits(finish)
	upperDigits2 := getDigits(start - 1)
	var num int64
	n := len(s)
	var base int64 = 1
	for i := 0; i < n; i++ {
		num += int64(s[n-1-i]-48) * base
		base *= 10
	}
	digits := getDigits(num)
	c1 := count(upperDigits1, digits, 15, limit, n)
	c2 := count(upperDigits2, digits, 15, limit, n)
	return c1 - c2
}
