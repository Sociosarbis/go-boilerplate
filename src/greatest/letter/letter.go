package letter

func greatestLetter(s string) string {
	m := make([]byte, 26)
	for _, c := range s {
		if c >= 97 {
			if m[c-97]&2 == 0 {
				m[c-97] |= 2
			}
		} else {
			if m[c-65]&1 == 0 {
				m[c-65] |= 1
			}
		}
	}

	for i := 25; i >= 0; i-- {
		if m[i] == 3 {
			return string([]byte{byte(i + 65)})
		}
	}

	return ""
}
