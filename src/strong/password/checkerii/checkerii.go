package checkerii

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	mask := 0
	for i, c := range password {
		if i > 0 && password[i] == password[i-1] {
			return false
		}
		if c >= 97 && c <= 122 {
			if mask&1 == 0 {
				mask |= 1
			}
		} else if c >= 65 && c <= 90 {
			if mask&2 == 0 {
				mask |= 2
			}
		} else if c >= 48 && c <= 57 {
			if mask&4 == 0 {
				mask |= 4
			}
		} else {
			if mask&8 == 0 {
				mask |= 8
			}
		}
	}
	return mask&15 == 15
}
