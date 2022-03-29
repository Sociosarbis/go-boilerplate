package answers

func helper(answerKey string, k int, s rune) int {
	ret := 0
	temp := 0
	replacedPos := []int{}
	for i, ans := range answerKey {
		if ans != s {
			if len(replacedPos) < k {
				replacedPos = append(replacedPos, i)
			} else {
				temp = i - replacedPos[0] - 1
				replacedPos = append(replacedPos[1:], i)
			}
		}
		temp += 1
		if temp > ret {
			ret = temp
		}
	}
	return ret
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	if k >= len(answerKey)>>1 {
		return len(answerKey)
	}
	opt1 := helper(answerKey, k, 'T')
	opt2 := helper(answerKey, k, 'F')
	if opt1 > opt2 {
		return opt1
	}
	return opt2
}
