package multiply2

const mask int = 1e9 + 7

func qpow(n int64, k int64) int64 {
	var out int64 = 1
	for k != 0 {
		if k&1 == 1 {
			out = (out * n) % int64(mask)
			k--
		} else {
			n = (n * n) % int64(mask)
			k >>= 1
		}
	}
	return out
}

func sumAndMultiply(s string, queries [][]int) []int {
	n := len(s)
	prefixSum := make([]int, n+1)
	prefixX := make([]int, n+1)
	prefixPower10 := make([]int, n+1)
	base := 1
	for i := n - 1; i >= 0; i-- {
		if s[i] != '0' {
			prefixX[i] = int((int64(prefixX[i+1]) + int64(s[i]-'0')*int64(base)) % int64(mask))
			base = int(int64(base) * 10 % int64(mask))
		} else {
			prefixX[i] = prefixX[i+1]
		}
	}
	for i := 0; i < n; i++ {
		if s[i] != '0' {
			prefixSum[i+1] = (prefixSum[i] + int(s[i]-'0')) % mask
			prefixPower10[i+1] = prefixPower10[i] + 1
		} else {
			prefixSum[i+1] = prefixSum[i]
			prefixPower10[i+1] = prefixPower10[i]
		}
	}

	out := make([]int, len(queries))
	for i, query := range queries {
		x := prefixX[query[0]] - prefixX[query[1]+1]
		if x < 0 {
			x += mask
		}
		sum := prefixSum[query[1]+1] - prefixSum[query[0]]
		if sum < 0 {
			sum += mask
		}
		x = int(int64(x) * qpow(10, int64(mask-2)*int64(prefixPower10[n]-prefixPower10[query[1]+1])) % int64(mask))
		out[i] = int(int64(x) * int64(sum) % int64(mask))
	}
	return out
}
