package chars2

func numberOfSpecialChars(word string) int {
	set := [26]byte{}
	var out int
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			index := c - 'A'
			if set[index] == 1 {
				out++
				set[index] = 2
			} else if set[index] == 0 {
				set[index] = 3
			}
		} else {
			index := c - 'a'
			if set[index] == 0 {
				set[index] = 1
			} else if set[index] == 2 {
				out--
				set[index] = 3
			}
		}
	}
	return out
}
