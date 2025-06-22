package string

func divideString(s string, k int, fill byte) []string {
	l := len(s)
	n := l / k
	out := make([]string, 0, n+1)
	for i := 0; i < n; i++ {
		out = append(out, s[i*k:(i+1)*k])
	}
	if n*k != l {
		temp := make([]byte, k)
		copy(temp, s[n*k:l])
		for i := l - n*k; i < k; i++ {
			temp[i] = fill
		}
		out = append(out, string(temp))
	}
	return out
}
