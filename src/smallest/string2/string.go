package string2

func smallestString(s string) string {
	bs := []byte(s)
	var hasStart bool
	for i, b := range bs {
		if hasStart {
			if b == 97 {
				break
			} else {
				bs[i] -= 1
			}
		} else {
			if b != 97 {
				hasStart = true
				bs[i] -= 1
			}
		}
	}
	if !hasStart {
		bs[len(bs)-1] = 'z'
	}
	return string(bs)
}
