package sum

func sumDigits(s []byte, i, j, to int) int {
	var sum int
	for k := j; k >= i; k-- {
		sum += int(s[k] - '0')
	}
	if sum == 0 {
		s[to] = '0'
		return 1
	}
	k := to
	var out int
	for sum != 0 {
		s[k] = byte(sum%10) + '0'
		sum /= 10
		k++
		out++
	}
	l := to
	r := to + out - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return out
}

func digitSum(s string, k int) string {
	n := len(s)
	b := []byte(s)
	for n > k {
		nn := 0
		for i := 0; i < n; i += k {
			end := i + k
			if end > n {
				end = n
			}
			nn += sumDigits(b, i, end-1, nn)
		}
		n = nn
	}
	return string(b[:n])
}
