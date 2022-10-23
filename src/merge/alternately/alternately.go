package alternately

func mergeAlternately(word1 string, word2 string) string {
	ret := make([]byte, len(word1)+len(word2))
	for i := 0; i < len(ret); i++ {
		if i%2 == 0 {
			if i/2 < len(word1) {
				ret[i] = word1[i/2]
			} else {
				copy(ret[i:], word2[len(word2)-(len(ret)-i):])
				break
			}
		} else {
			if i/2 < len(word2) {
				ret[i] = word2[i/2]
			} else {
				copy(ret[i:], word1[len(word1)-(len(ret)-i):])
			}
		}
	}
	return string(ret)
}
