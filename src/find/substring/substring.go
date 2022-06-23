package substring

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	wordMap := map[string]int{}
	for _, word := range words {
		count := wordMap[word]
		wordMap[word] = count + 1
	}

	ret := []int{}

	for i := 0; i < wordLen; i++ {
		tempMap := map[string]int{}
		diff := len(wordMap)
		rightBound := i + wordLen*len(words)
		if rightBound > len(s) {
			break
		}
		for j := rightBound - wordLen*len(words); j < rightBound; j += wordLen {
			w := s[j : j+wordLen]
			count := tempMap[w]
			count++
			tempMap[w] = count
			count2 := wordMap[w]
			if count2 == count {
				diff--
			} else if count2+1 == count {
				diff++
			}
		}
		if diff == 0 {
			ret = append(ret, rightBound-wordLen*len(words))
		}
		rightBound += wordLen
		for rightBound <= len(s) {
			prevW := s[rightBound-wordLen*(len(words)+1) : rightBound-wordLen*len(words)]
			w := s[rightBound-wordLen : rightBound]
			count := tempMap[prevW]
			count2 := wordMap[prevW]
			if count2+1 == count {
				diff--
			} else if count2 == count {
				diff++
			}
			tempMap[prevW] = count - 1
			count = tempMap[w]
			count++
			tempMap[w] = count
			count2 = wordMap[w]
			if count2 == count {
				diff--
			} else if count2+1 == count {
				diff++
			}
			if diff == 0 {
				ret = append(ret, rightBound-wordLen*len(words))
			}
			rightBound += wordLen
		}
	}
	return ret
}
