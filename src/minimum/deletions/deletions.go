package deletions

import "sort"

func minimumDeletions(word string, k int) int {
	counter := make([]int, 26)
	for _, c := range word {
		counter[c-'a']++
	}
	sort.Ints(counter)
	var i int
	ret := len(word)
	var deletions int
	for ; i < 26; i++ {
		if counter[i] != 0 {
			if i == 0 || (i > 0 && counter[i] != counter[i-1]) {
				var temp int
				for j := 25; j > i; j-- {
					if counter[i]+k < counter[j] {
						temp += counter[j] - counter[i] - k
					} else {
						break
					}
				}
				if temp+deletions < ret {
					ret = temp + deletions
				}
			}
			deletions += counter[i]
		}
	}
	return ret
}
