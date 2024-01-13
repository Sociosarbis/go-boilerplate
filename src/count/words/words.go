package words

func countWords(words1 []string, words2 []string) int {
	dict := make(map[string]int, len(words1))

	for _, s := range words1 {
		if count, ok := dict[s]; !ok || count < 2 {
			if !ok {
				dict[s] = 1
			} else {
				dict[s] = 2
			}
		}
	}

	ret := 0
	for _, s := range words2 {
		if count, ok := dict[s]; ok && count <= 1 && count > -1 {
			dict[s] -= 1
		}
	}

	for _, count := range dict {
		if count == 0 {
			ret++
		}
	}
	return ret
}
