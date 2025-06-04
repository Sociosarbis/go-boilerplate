package string

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	n := len(word)
	var start int
	maxLen := n - numFriends + 1
	l := maxLen
	for i := 1; i < n; i++ {
		end := i + maxLen
		if end > n {
			end = n
		}
		for j := i; j < end; j++ {
			if word[j] > word[start+j-i] {
				start, l = i, end-i
				break
			} else if word[j] < word[start+j-i] {
				break
			}
		}
	}
	return word[start : start+l]
}
