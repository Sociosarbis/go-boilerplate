package array

func divisibilityArray(word string, m int) []int {
	n := len(word)
	ret := make([]int, n)
	remain := 0
	for i, c := range word {
		temp := remain*10 + int(c-48)
		remain = temp % m
		if remain == 0 {
			ret[i] = 1
		}
	}
	return ret
}
