package number3

var primeNums = [4]int64{2, 3, 5, 7}

func gcd(a, b int64) int64 {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func smallestNumber(num string, t int64) string {
	temp := t
	for _, num := range primeNums {
		for temp > 1 && temp%num == 0 {
			temp /= num
		}
	}
	if temp != 1 {
		return "-1"
	}
	n := len(num)
	rem := make([]int64, n+1)
	rem[0] = t
	pos := n - 1
	for i, c := range num {
		if c == '0' {
			pos = i
			break
		}
		rem[i+1] = rem[i] / gcd(rem[i], int64(c-'0'))
	}
	if rem[n] == 1 {
		return num
	}
	digits := []byte(num)
	// 从第一个0位或者从最后一位开始
	for i := pos; i >= 0; i-- {
		digits[i]++
		for digits[i] <= '9' {
			tNow := rem[i] / gcd(rem[i], int64(digits[i]-'0'))
			var k int64 = 9
			// 从最大的数字，最后一位开始尝试填入数字
			for j := n - 1; j > i; j-- {
				for tNow%k != 0 {
					k--
				}
				tNow /= k
				digits[j] = '0' + byte(k)
			}
			if tNow == 1 {
				return string(digits)
			}
			digits[i]++
		}
	}
	// 如果在n位数内也无法使t等于1
	ans := make([]byte, 0, n+1)
	var i int64 = 9
	for ; i > 1; i-- {
		for t%i == 0 {
			ans = append(ans, '0'+byte(i))
			t /= i
		}
	}
	for len(ans) <= n {
		ans = append(ans, '1')
	}
	var l int
	r := len(ans) - 1
	for l < r {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}
	return string(ans)
}
