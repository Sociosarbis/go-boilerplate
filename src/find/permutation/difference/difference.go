package difference

func findPermutationDifference(s string, t string) int {
	n := len(s)
	indices := make([]int, 26)
	for i := 0; i < n; i++ {
		indices[s[i]-97] = i
	}
	var ret int
	for i := 0; i < n; i++ {
		temp := indices[t[i]-97] - i
		if temp < 0 {
			ret -= temp
		} else {
			ret += temp
		}
	}
	return ret
}
