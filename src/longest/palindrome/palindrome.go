package palindrome

func longestPalindrome(words []string) int {
	var counter [676]int
	var ret int
	for _, word := range words {
		b := int(word[1]-'a')*26 + int(word[0]-'a')
		if counter[b] > 0 {
			ret += 4
			counter[b]--
		} else {
			counter[int(word[0]-'a')*26+int(word[1]-'a')]++
		}
	}
	var temp int
	for i := 0; i < 26; i++ {
		if counter[i*26+i] > temp {
			temp = counter[i*26+i]
		}
	}
	ret += temp * 2
	return ret
}
