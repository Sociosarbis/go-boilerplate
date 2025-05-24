package containing

import "strings"

func findWordsContaining(words []string, x byte) []int {
	var ret []int
	for i, word := range words {
		if strings.IndexByte(word, x) != -1 {
			ret = append(ret, i)
		}
	}
	return ret
}
