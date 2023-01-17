package similar

func toWords(sentence string) []string {
	ret := []string{}
	i := 0
	for i < len(sentence) {
		j := i
		for ; j < len(sentence); j++ {
			if sentence[j] == ' ' {
				break
			}
		}
		ret = append(ret, sentence[i:j])
		i = j + 1
	}
	return ret
}

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1, words2 := toWords(sentence1), toWords(sentence2)
	if len(words1) > len(words2) {
		words1, words2 = words2, words1
	}
	diff := len(words2) - len(words1)
loop:
	for i := 0; i <= len(words1); i++ {
		j := 0
		for ; j < i; j++ {
			if words1[j] != words2[j] {
				continue loop
			}
		}
		for ; j < len(words1); j++ {
			if words1[j] != words2[j+diff] {
				continue loop
			}
		}
		return true
	}
	return false
}
