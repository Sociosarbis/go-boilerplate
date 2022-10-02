package transform

func canTransform(start string, end string) bool {
	n := len(start)
	s := []byte(start)
loop:
	for i := 0; i < n; i++ {
		a := s[i]
		b := end[i]
		if a == 'R' {
			if b == 'L' {
				return false
			} else if b == 'X' {
				for j := i + 1; j < n; j++ {
					if s[j] == 'X' {
						s[i], s[j] = s[j], s[i]
						continue loop
					} else if s[j] == 'L' {
						return false
					}
				}
				return false
			}
		} else if a == 'X' {
			if b == 'R' {
				return false
			} else if b == 'L' {
				for j := i + 1; j < n; j++ {
					if s[j] == 'L' {
						s[i], s[j] = s[j], s[i]
						continue loop
					} else if s[j] == 'R' {
						return false
					}
				}
				return false
			}
		} else {
			if b != 'L' {
				return false
			}
		}
	}
	return true
}
