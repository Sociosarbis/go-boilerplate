package palindrome

func smallestPalindrome(s string) string {
	counter := [26]int{}
	for _, c := range s {
		counter[c-'a']++
	}
	out := make([]byte, 0, len(s))
	oddIndex := -1
	for i := 0; i < 26; i++ {
		if counter[i] == 0 {
			continue
		}
		c := 'a' + byte(i)
		n := counter[i] / 2
		for j := 0; j < n; j++ {
			out = append(out, c)
		}
		if counter[i]%2 != 0 {
			oddIndex = i
		}
	}
	var r int
	if oddIndex != -1 {
		c := 'a' + byte(oddIndex)
		out = append(out, c)
		r = len(out) - 2
	} else {
		r = len(out) - 1
	}
	for i := r; i >= 0; i-- {
		out = append(out, out[i])
	}
	return string(out)
}
