package number

func reformatNumber(number string) string {
	count := 0
	for _, c := range number {
		if c >= 48 && c <= 57 {
			count++
		}
	}
	temp := count % 3
	l := count
	if temp == 0 {
		l += count/3 - 1
	} else {
		l += count / 3
	}
	ret := make([]byte, l)
	i := 0
	count2 := 0
	for _, c := range number {
		if c >= 48 && c <= 57 {
			ret[i] = byte(c)
			count2++
			i++
			if temp == 1 {
				if count2+4 > count {
					if count2+2 == count {
						ret[i] = '-'
						i++
					}
					continue
				}
			}
			if count2 != count && count2%3 == 0 {
				ret[i] = '-'
				i++
			}
		}
	}
	return string(ret)
}
