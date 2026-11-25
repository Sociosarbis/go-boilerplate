package subarray

const base int = 131
const mod int = 1e9 + 7

func ensurePositive(p int) int {
	if p < 0 {
		return p + mod
	}
	return p
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % int64(mod))
}

func smallestUniqueSubarray(nums []int) int {
	n := len(nums)
	pow := make([]int, n+1)
	pow[0] = 1
	for i := 1; i <= n; i++ {
		pow[i] = mul(pow[i-1], base)
	}
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = (mul(prefix[i-1], base) + nums[i-1]) % mod
	}
	l, r := 1, n
	out := r
loop:
	for l <= r {
		mid := (l + r) >> 1
		m := make(map[int]byte, n-mid+1)
		p := pow[mid]
		for i := mid - 1; i < n; i++ {
			h := ensurePositive(prefix[i+1] - mul(prefix[i-mid+1], p))
			if c, ok := m[h]; ok {
				if c < 2 {
					m[h] = c + 1
				}
			} else {
				m[h] = 1
			}
		}
		for _, c := range m {
			if c == 1 {
				out = mid
				r = mid - 1
				continue loop
			}
		}
		l = mid + 1
	}
	return out
}
