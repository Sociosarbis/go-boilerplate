package char

func shortestToChar(s string, c byte) []int {
	ret := make([]int, len(s))
	temp := len(s)
	for i := 0; i < len(s); i++ {
		if s[i] != c {
			temp++
			ret[i] = temp
		} else {
			temp = 1
			for j := i - 1; j >= 0; j-- {
				if ret[j] > temp {
					ret[j] = temp
					temp++
				} else {
					break
				}
			}
			temp = 0
		}
	}
	return ret
}
