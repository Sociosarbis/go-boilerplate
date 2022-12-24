package merge

func largestMerge(word1 string, word2 string) string {
	i := 0
	j := 0
	index := 0
	ret := make([]byte, len(word1)+len(word2))
	for index < len(ret) {
		if i < len(word1) {
			if j < len(word2) {
				choose1 := false
				if word1[i] > word2[j] {
					choose1 = true
				} else if word1[i] == word2[j] {
					k := 1
					for i+k < len(word1) && j+k < len(word2) {
						if word1[i+k] > word2[j+k] {
							choose1 = true
							break
						} else if word1[i+k] < word2[j+k] {
							break
						}
						k++
					}
					if i+k < len(word1) && j+k == len(word2) {
						choose1 = true
					}
				}
				if choose1 {
					ret[index] = word1[i]
					i++
					index++
				} else {
					ret[index] = word2[j]
					j++
					index++
				}
			} else {
				copy(ret[index:], word1[i:])
				break
			}
		} else {
			copy(ret[index:], word2[j:])
			break
		}
	}
	return string(ret)
}
