package string

func smallestBeautifulString(s string, k int) string {
	n := len(s)
	bs := []byte(s)
	maxB := byte(97 + k - 1)
	i := n - 1
loop1:
	for ; i >= 0; i-- {
		b := s[i] + 1
		for b <= maxB {
			if canPlace(bs, i, b) {
				bs[i] = b
				break loop1
			}
			b++
		}
	}
	if i == -1 {
		return ""
	}
	from := byte(97)
loop2:
	for j := i + 1; j < n; j++ {
		b := from
		for b <= maxB {
			if canPlace(bs, j, b) {
				bs[j] = b
				continue loop2
			}
			b++
		}
		return ""
	}
	return string(bs)
}

func canPlace(s []byte, i int, c byte) bool {
	if i > 0 {
		if s[i-1] != c {
			if i > 1 {
				if s[i-2] != c {
					return true
				}
			} else {
				return true
			}
		}
		return false
	}
	return true
}
