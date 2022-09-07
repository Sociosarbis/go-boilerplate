package spaces

func reorderSpaces(text string) string {
	words := []string{}
	spaces := 0
	i := 0
	for i < len(text) {
		start := i
		for i < len(text) && text[i] != ' ' {
			i++
		}
		if i == start {
			i++
			spaces++
		} else {
			words = append(words, text[start:i])
		}
	}
	var interval int
	var rest int
	if len(words) > 1 {
		interval = spaces / (len(words) - 1)
		rest = spaces % (len(words) - 1)
	} else {
		rest = spaces
	}

	ret := make([]byte, len(text))
	index := 0
	for i := 0; i < len(words); i++ {
		copy(ret[index:], []byte(words[i]))
		index += len(words[i])

		if i+1 < len(words) {
			spaces = interval
		} else {
			spaces = rest
		}
		for j := 0; j < spaces; j++ {
			ret[index] = ' '
			index++
		}
	}
	return string(ret)
}
