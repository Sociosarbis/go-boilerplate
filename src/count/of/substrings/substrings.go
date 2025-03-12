package substrings

var mappings = [26]int{}

func init() {
	mappings[0], mappings['e'-'a'], mappings['i'-'a'], mappings['o'-'a'], mappings['u'-'a'] = 1, 2, 3, 4, 5
}

func getIndex(c byte) int {
	index := mappings[c-'a']
	if index > 0 {
		return index - 1
	} else {
		return 5
	}
}

func isMatch(counter *[6]int, k int) bool {
	if counter[5] == k {
		for i := 0; i < 5; i++ {
			if counter[i] == 0 {
				return false
			}
		}
		return true
	}
	return false
}

func countOfSubstrings(word string, k int) int {
	c1 := [6]int{}
	start := k + 5
	for i := 0; i < start; i++ {
		c1[getIndex(word[i])]++
	}
	n := len(word)
	var ret int
	for i := start; i <= n; i++ {
		if isMatch(&c1, k) {
			ret++
		}
		temp := c1
		for j := i; j < n; j++ {
			temp[getIndex(word[j-i])]--
			temp[getIndex(word[j])]++
			if isMatch(&temp, k) {
				ret++
			}
		}
		if i < n {
			c1[getIndex(word[i])]++
		}
	}
	return ret
}
