package ascending

import "strconv"

func areNumbersAscending(s string) bool {
	i := 0
	var prev int64 = -1
	for i < len(s) {
		j := 0
		for ; i+j < len(s) && s[i+j] != ' '; j++ {
		}
		num, err := strconv.ParseInt(s[i:i+j], 10, 32)
		if err == nil {
			if num > prev {
				prev = num
			} else {
				return false
			}
		}
		i += j + 1
	}
	return true
}
