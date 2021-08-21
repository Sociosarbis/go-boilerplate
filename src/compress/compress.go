package compress

func compress(chars []byte) int {
	l := len(chars) + 1
	count := 1
	k := 0
	cur := chars[0]
	for i := 1; i < l; i += 1 {
		if i < len(chars) && chars[i] == cur {
			count += 1
		} else {
			j := 0
			temp := count
			if temp > 1 {
				for temp > 0 {
					j += 1
					temp /= 10
				}
			}
			chars[k] = cur
			for ii := j; ii > 0; ii -= 1 {
				chars[k+ii] = byte(count%10 + 48)
				count /= 10
			}
			k += j + 1
			if i < len(chars) {
				count = 1
				cur = chars[i]
			}
		}
	}
	return k
}
