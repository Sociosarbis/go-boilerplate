package string

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	groups := [26]byte{}
	for i := 0; i < 26; i++ {
		groups[i] = byte(i + 'a')
	}
	n := len(s1)

	for i := 0; i < n; i++ {
		g1, g2 := groups[s1[i]-'a'], groups[s2[i]-'a']
		if g1 != g2 {
			if g1 > g2 {
				g1, g2 = g2, g1
			}
			for i := 0; i < 26; i++ {
				if groups[i] == g2 {
					groups[i] = g1
				}
			}
		}
	}
	out := []byte(baseStr)
	for i, c := range out {
		if groups[c-'a'] != c {
			out[i] = groups[c-'a']
		}
	}
	return string(out)
}
