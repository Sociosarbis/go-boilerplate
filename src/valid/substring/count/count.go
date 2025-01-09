package count

func isValid(counter []int, target []int) bool {
	for i, c := range counter {
		if c < target[i] {
			return false
		}
	}
	return true
}

func validSubstringCount(word1 string, word2 string) int64 {
	target := make([]int, 26)
	for _, c := range word2 {
		target[c-97]++
	}
	var ret int64
	n := len(word1)
	counter := make([]int, 26)
	var j int
	for i := 0; i < n; i++ {
		if j < i {
			j = i
		}
		for j < n && !isValid(counter, target) {
			counter[word1[j]-97]++
			j++
		}
		if isValid(counter, target) {
			ret += int64(n - j + 1)
		}
		counter[word1[i]-97]--
	}
	return ret
}
