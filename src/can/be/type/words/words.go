package words

func canBeTypedWords(text string, brokenLetters string) int {
	var hasLetter bool
	var broken bool
	brokenLetterMap := [26]bool{}
	for i := range brokenLetters {
		brokenLetterMap[brokenLetters[i]-'a'] = true
	}
	var out int
	for i := range text {
		if hasLetter {
			if text[i] == ' ' {
				if !broken {
					out++
				}
				hasLetter = false
			} else {
				if !broken && brokenLetterMap[text[i]-'a'] {
					broken = true
				}
			}
		} else {
			if brokenLetterMap[text[i]-'a'] {
				broken = true
			} else {
				broken = false
			}
			hasLetter = true
		}
	}
	if !broken {
		return out + 1
	}
	return out
}
