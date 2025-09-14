package checker

import "strings"

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func replaceVowel(s string) string {
	temp := []byte(s)
	n := len(temp)
	for i := 0; i < n; i++ {
		if isVowel(temp[i]) {
			temp[i] = 'a'
		}
	}
	return string(temp)
}

func spellchecker(wordlist []string, queries []string) []string {
	n := len(wordlist)
	ms := [3]map[string]int{}
	for i := 0; i < 3; i++ {
		ms[i] = make(map[string]int, n)
	}
	for i, word := range wordlist {
		if _, ok := ms[0][word]; !ok {
			ms[0][word] = i
		}
		temp := strings.ToLower(word)
		if _, ok := ms[1][temp]; !ok {
			ms[1][temp] = i
		}
		temp, word = replaceVowel(temp), temp
		if _, ok := ms[2][temp]; !ok {
			ms[2][temp] = i
		}
	}
	out := make([]string, len(queries))
	for i, query := range queries {
		if j, ok := ms[0][query]; ok {
			out[i] = wordlist[j]
			continue
		}
		query = strings.ToLower(query)
		if j, ok := ms[1][query]; ok {
			out[i] = wordlist[j]
			continue
		}
		query = replaceVowel(query)
		if j, ok := ms[2][query]; ok {
			out[i] = wordlist[j]
			continue
		}
	}
	return out
}
