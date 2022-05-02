package valid

func isValid(code string) bool {
	if len(code) < 7 {
		return false
	}
	tagStack := []string{}
	i := 0
	if code[i] != '<' {
		return false
	}
	i++
	j := i
	for j < len(code) && code[j] >= 65 && code[j] <= 90 {
		j++
	}
	if j-i > 9 || j-i < 1 || j >= len(code) || code[j] != '>' {
		return false
	}
	tagStack = append(tagStack, code[i:j])
	j++

	isCDATA := false

	for j < len(code) {
		if !isCDATA {
			if code[j] == '<' {
				if j+1 < len(code) {
					if code[j+1] == '/' {
						j += 2
						i = j
						for j < len(code) && code[j] >= 65 && code[j] <= 90 {
							j++
						}
						if j-i > 9 || j-i < 1 || j >= len(code) || code[j] != '>' {
							return false
						}
						if code[i:j] == tagStack[len(tagStack)-1] {
							tagStack = tagStack[:len(tagStack)-1]
							if len(tagStack) == 0 && j+1 != len(code) {
								return false
							}
						} else {
							return false
						}
						j++
					} else if j+8 < len(code) && code[j+1:j+9] == "![CDATA[" {
						isCDATA = true
						j += 9
					} else if code[j+1] >= 65 && code[j+1] <= 90 {
						j++
						i = j
						for j < len(code) && code[j] >= 65 && code[j] <= 90 {
							j++
						}
						if j-i > 9 || j-i < 1 || j >= len(code) || code[j] != '>' {
							return false
						}
						tagStack = append(tagStack, code[i:j])
						j++
					} else {
						return false
					}
				}
			} else {
				j++
			}
		} else {
			i = j
			for j < len(code) {
				if j-i >= 2 {
					if code[j-2:j+1] == "]]>" {
						j++
						break
					}
				}
				j++
			}
			isCDATA = false
		}
	}

	return len(tagStack) == 0
}
