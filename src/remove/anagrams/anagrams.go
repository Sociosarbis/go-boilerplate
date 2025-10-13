package anagrams

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	counter := [26]int{}
	for i := 0; i < len(a); i++ {
		counter[a[i]-'a']++
	}
	for i := 0; i < len(b); i++ {
		index := b[i] - 'a'
		counter[index]--
		if counter[index] < 0 {
			return false
		}
	}
	return true
}

func removeAnagrams(words []string) []string {
	i := 1
	n := len(words)
	for j := i; j < n; j++ {
		if !isAnagram(words[i-1], words[j]) {
			words[i] = words[j]
			i++
		}
	}
	return words[:i]
}
