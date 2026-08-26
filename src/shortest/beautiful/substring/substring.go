package substring

func isSmall(a, b string) bool {
	n1, n2 := len(a), len(b)
	if n1 < n2 {
		return true
	} else if n1 > n2 {
		return false
	}
	for i := 0; i < n1; i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return false
}

func shortestBeautifulSubstring(s string, k int) string {
	var count int
	n := len(s)
	var l int
	for ; l < n; l++ {
		if s[l] == '1' {
			count++
			break
		}
	}
	if count == k {
		return "1"
	}
	var out string
	for i := l + 1; i < n; i++ {
		if s[i] == '1' {
			count++
			if count == k {
				if len(out) == 0 {
					out = s[l : i+1]
				} else if isSmall(s[l:i+1], out) {
					out = s[l : i+1]
				}
				l++
				for ; l < n; l++ {
					if s[l] == '1' {
						break
					}
				}
				count--
			}
		}
	}
	return out
}
