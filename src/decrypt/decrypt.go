package decrypt

func decrypt(code []int, k int) []int {
	ret := make([]int, len(code))
	temp := 0
	index := 0
	if k > 0 {
		for i := 1; i <= k; i++ {
			index = i % len(code)
			temp += code[index]
		}
	} else if k < 0 {
		for i := -1; i >= k; i-- {
			index = (len(code) + i) % len(code)
			temp += code[index]
		}
	}
	ret[0] = temp
	for i := 1; i < len(code); i++ {
		if k > 0 {
			index = (index + 1) % len(code)
			temp += code[index] - code[i]
		} else if k < 0 {
			temp -= code[index]
			index = (index + 1) % len(code)
			temp += code[i-1]
		}
		ret[i] = temp
	}
	return ret
}
