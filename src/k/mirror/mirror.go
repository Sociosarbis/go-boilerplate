package mirror

func getNum(digits []int, n int) int64 {
	var base int64 = 1
	var temp int64
	half := n / 2
	for i := n - 1; i >= 0; i-- {
		if i >= half {
			temp += int64(digits[n-1-i]) * base
		} else {
			temp += int64(digits[i]) * base
		}
		base *= 10
	}
	return temp
}

func isBaseK(num, k int64) bool {
	digits := []int{}
	for num != 0 {
		digits = append(digits, int(num%k))
		num -= num % k
		num /= k
	}
	var l int
	r := len(digits) - 1
	for l < r {
		if digits[l] != digits[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func iter(digits []int, i, k, m int, n *int) int64 {
	if *n == 0 {
		return 0
	}
	var ret int64
	if i < len(digits) {
		var start int
		if i == 0 {
			start = 1
		}
		for j := start; j <= 9; j++ {
			digits[i] = j
			ret += iter(digits, i+1, k, m, n)
		}
	} else {
		d := getNum(digits, m)
		if isBaseK(d, int64(k)) {
			*n--
			ret += d
		}
	}
	return ret
}

func kMirror(k int, n int) int64 {
	i := 1
	count := n
	var out int64
	for count != 0 {
		digits := make([]int, i)
		out += iter(digits, 0, k, i*2-1, &count)
		if count != 0 {
			out += iter(digits, 0, k, i*2, &count)
		}
		i++
	}
	return out
}
