package count

func digitCount(num string) bool {
	temp := make([]byte, len(num))
	for _, c := range num {
		if int(c-48) < len(num) {
			temp[c-48]++
		} else {
			return false
		}
	}
	for i := range temp {
		if temp[i]+48 != num[i] {
			return false
		}
	}
	return true
}
