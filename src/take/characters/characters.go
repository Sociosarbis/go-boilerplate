package characters

func isPassed(counter [3]int, k int) bool {
	for i := 0; i < 3; i++ {
		if counter[i] < k {
			return false
		}
	}
	return true
}

func takeCharacters(s string, k int) int {
	n := len(s)
	counter := [3]int{}
	for i := 0; i < n; i++ {
		counter[s[i]-'a']++
	}
	if !isPassed(counter, k) {
		return -1
	}
	ret := n
	j := n
	for i := n - 1; i >= 0; i-- {
		counter[s[i]-'a']--
		for j > i && !isPassed(counter, k) {
			counter[s[j-1]-'a']++
			j--
		}
		if n-j+i < ret {
			ret = n - j + i
		}
	}
	return ret
}
