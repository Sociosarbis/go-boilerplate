package arrays

const mask int64 = 1e9 + 7

func mulRange(l, r int64) int64 {
	var out int64 = 1
	for l <= r {
		out = (out * l) % mask
		l++
	}
	return out
}

func qpow(n int64, k int) int64 {
	var out int64 = 1
	for k != 0 {
		if k&1 == 1 {
			out = (out * n) % mask
			k--
		} else {
			n = (n * n) % mask
			k >>= 1
		}
	}
	return out
}

func countGoodArrays(n int, m int, k int) int {
	// 从n - 1个空位，保留k个空位置，其他插隔板
	// 根据费马小定理，(a / b) % m = a * b ^ (m - 2) % m
	return int((((int64(m) * mulRange(int64(n-k), int64(n-1))) % mask) * qpow(mulRange(2, int64(k)), int(mask-2)) % mask) * qpow(int64(m-1), n-1-k) % mask)
}
