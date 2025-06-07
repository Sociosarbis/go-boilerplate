package stars

func clearStars(s string) string {
	n := len(s)
	indices := [26][]int{}
	removed := make([]bool, n)
	l := n
	for i := 0; i < n; i++ {
		if s[i] == '*' {
			for j := 0; j < 26; j++ {
				if len(indices[j]) != 0 {
					removed[indices[j][len(indices[j])-1]] = true
					indices[j] = indices[j][:len(indices[j])-1]
					break
				}
			}
			l -= 2
		} else {
			indices[s[i]-'a'] = append(indices[s[i]-'a'], i)
		}
	}
	out := make([]byte, 0, l)
	for i := 0; i < n; i++ {
		if s[i] != '*' && !removed[i] {
			out = append(out, s[i])
		}
	}
	return string(out)
}
