package latin

func toGoatLatin(sentence string) string {
	i := 0
	ret := ""
	j := 0
	for j < len(sentence) {
		if sentence[j] != ' ' {
			i++
			l := j
			for j < len(sentence) && sentence[j] != ' ' {
				j++
			}
			isVowel := false
			firstChar := sentence[l]
			if firstChar >= 65 && firstChar <= 90 {
				firstChar += 32
			}
			if firstChar == 'a' || firstChar == 'e' || firstChar == 'i' || firstChar == 'o' || firstChar == 'u' {
				isVowel = true
			}
			part := make([]byte, j-l+2+i)
			if isVowel {
				copy(part, []byte(sentence)[l:j])
			} else {
				copy(part, []byte(sentence)[l+1:j])
				part[j-l-1] = sentence[l]
			}
			part[j-l] = 'm'
			for k := j - l + 1; k < len(part); k++ {
				part[k] = 'a'
			}
			ret += string(part)
		} else {
			ret += " "
			j++
		}
	}
	return ret
}
