package swaps

func minSwaps(s string) int {
	n := len(s)
	var temp int
	j := n
	var ret int
	for i := 0; i < j; i++ {
		if s[i] == '[' {
			temp++
		} else {
			temp--
		}
		if temp < 0 {
			for j-1 > i {
				if s[j-1] == '[' {
					ret++
					j--
					temp = 1
					break
				}
				j--
			}
		}
	}
	return ret
}
