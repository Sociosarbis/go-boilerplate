package substring

func maximumLengthSubstring(s string) int {
	counter := [26]int{}
	var i int
	var out int
	for j, c := range s {
		index := c - 'a'
		counter[index]++
		for counter[index] > 2 {
			counter[s[i]-'a']--
			i++
		}
		if j-i+1 > out {
			out = j - i + 1
		}
	}
	return out
}
