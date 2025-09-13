package sum

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func maxFreqSum(s string) int {
	counter := [26]int{}
	n := len(s)
	for i := 0; i < n; i++ {
		counter[s[i]-'a']++
	}
	var a int
	var b int
	var i byte
	for ; i < 26; i++ {
		if counter[i] > a || counter[i] > b {
			if isVowel(i + 'a') {
				if counter[i] > a {
					a = counter[i]
				}
			} else {
				if counter[i] > b {
					b = counter[i]
				}
			}
		}
	}
	return a + b
}
