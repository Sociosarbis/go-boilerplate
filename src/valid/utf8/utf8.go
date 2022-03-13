package utf8

func validUtf8(data []int) bool {
	byteCount := 0
	for _, num := range data {
		if byteCount == 0 {
			mask := 1 << 7
			for {
				if num&mask != 0 {
					byteCount += 1
					if byteCount > 4 {
						return false
					}
					mask >>= 1
				} else {
					break
				}
			}
			if byteCount > 1 {
				byteCount -= 1
			} else if byteCount == 1 {
				return false
			}
		} else {
			if num&128 != 0 && num&64 == 0 {
				byteCount -= 1
			} else {
				return false
			}
		}
	}
	return byteCount == 0
}
