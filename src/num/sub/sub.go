package sub

const mod int64 = 1e9 + 7

func numSub(s string) int {
	var temp int
	var out int64
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			temp++
		} else {
			out = (out + (int64(1+temp)*int64(temp)/2)%mod) % mod
			temp = 0
		}
	}
	if temp > 0 {
		out = (out + (int64(1+temp)*int64(temp)/2)%mod) % mod
	}
	return int(out)
}
