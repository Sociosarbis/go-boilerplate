package length

func minAnagramLength(s string) int {
	n := len(s)
	c1 := make([]int, 26)
loop:
	for i := 1; i <= n; i++ {
		c1[s[i-1]-97]++
		if n%i == 0 {
			for j := i; j < n; j += i {
				c2 := make([]int, 26)
				for k := 0; k < i; k++ {
					c2[s[j+k]-97]++
				}
				for k := 0; k < 26; k++ {
					if c1[k] != c2[k] {
						continue loop
					}
				}
			}
			return i
		}
	}
	return n
}
