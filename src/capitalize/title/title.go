package title

func toLowerCase(c byte) byte {
	if c < 97 {
		return c + 32
	}
	return c
}

func toUpperCase(c byte) byte {
	if c >= 97 {
		return c - 32
	}
	return c
}

func capitalizeTitle(title string) string {
	temp := []byte(title)
	temp = append(temp, ' ')
	length := 0
	n := len(temp)
	for i := 0; i < n; i++ {
		if temp[i] == ' ' {
			var j int
			if length > 2 {
				temp[i-length] = toUpperCase(temp[i-length])
				j = i - length + 1
			} else {
				j = i - length
			}
			length = 0
			for ; j < i; j++ {
				temp[j] = toLowerCase(temp[j])
			}
		} else {
			length++
		}
	}
	return string(temp[:len(title)])
}
