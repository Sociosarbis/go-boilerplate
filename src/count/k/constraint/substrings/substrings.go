package substrings

func countKConstraintSubstrings(s string, k int) int {
	n := len(s)
	var ret int
	for i := 0; i < n; i++ {
		var a int
		var b int
		for j := i; j < n; j++ {
			if s[j] == '0' {
				a++
			} else {
				b++
			}
			if a <= k || b <= k {
				ret++
			}
		}
	}
	return ret
}
