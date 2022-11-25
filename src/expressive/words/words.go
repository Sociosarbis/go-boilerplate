package words

type group struct {
	ch byte
	c  int
}

func expressiveWords(s string, words []string) int {
	groups := []group{}
	for i := 0; i < len(s); i++ {
		j := i
		for ; j+1 < len(s); j++ {
			if s[j+1] != s[i] {
				break
			}
		}
		groups = append(groups, group{
			s[i],
			j - i + 1,
		})
		i = j
	}
	ret := 0
loop:
	for _, word := range words {
		index := 0
		for i := 0; i < len(word); i++ {
			j := i
			for ; j+1 < len(word); j++ {
				if word[j+1] != word[i] {
					break
				}
			}
			if index < len(groups) && word[i] == groups[index].ch {
				if (j-i+1 < groups[index].c && groups[index].c >= 3) || (j-i+1 == groups[index].c) {
					index++
					i = j
					continue
				}
			}
			continue loop
		}
		if index == len(groups) {
			ret++
		}
	}
	return ret
}
