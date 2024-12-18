package strings2

const MOD int64 = 1e9
const BASE int64 = 233

// 滚动哈希
type HashString struct {
	P []int64
	H []int64
}

func NewHashString(s string) HashString {
	n := len(s)
	P := make([]int64, n+1)
	P[0] = 1
	for i := 1; i <= n; i++ {
		P[i] = P[i-1] * BASE % MOD
	}
	H := make([]int64, n+1)
	H[0] = 0
	for i := 1; i <= n; i++ {
		H[i] = (H[i-1]*BASE + int64(s[i-1])) % MOD
	}
	return HashString{P, H}
}

func (hs *HashString) query(l, r int) int64 {
	return (hs.H[r] - hs.H[l-1]*hs.P[r-l+1]%MOD + MOD) % MOD
}

type empty struct {
}

func minValidStrings(words []string, target string) int {
	n := len(target)
	st := make([]map[int64]empty, n+1)
	for i := 0; i <= n; i++ {
		st[i] = map[int64]empty{}
	}
	for _, word := range words {
		hs := NewHashString(word)
		for i := 1; i <= n && i <= len(word); i++ {
			hash := hs.query(1, i)
			if _, ok := st[i][hash]; !ok {
				st[i][hash] = empty{}
			}
		}
	}
	hs := NewHashString(target)
	f := make([]int, n+1)
	var j int
	for i := 0; i < n; i++ {
		if i > j {
			break
		}

		if _, ok := st[1][hs.query(i+1, i+1)]; !ok {
			continue
		}
		head := i + 1
		tail := n
		for head < tail {
			mid := (head + tail + 1) >> 1
			if _, ok := st[mid-i][hs.query(i+1, mid)]; ok {
				head = mid
			} else {
				tail = mid - 1
			}
		}
		// f[i] <= f[j], for i < j
		for j < head {
			j++
			f[j] = f[i] + 1
		}
	}
	if j == n {
		return f[n]
	}
	return -1
}
