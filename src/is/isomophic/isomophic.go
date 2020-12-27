package isomophic

func isIsomorphic(s string, t string) bool {
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := range s {
		c, prs := s2t[s[i]]
		if prs {
			if c != t[i] {
				return false
			}
		} else {
			_, prs := t2s[t[i]]
			if prs {
				return false
			} else {
				s2t[s[i]] = t[i]
				t2s[t[i]] = s[i]
			}
		}
	}
	return true
}
