package vowels

import "sort"

func isVowel(b byte) bool {
	if b < 97 {
		b += 32
	}
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func sortVowels(s string) string {
	n := len(s)
	v := make([]byte, n)
	copy(v, s)
	sort.Slice(v, func(i, j int) bool {
		return v[i] < v[j]
	})
	out := []byte(s)
	var j int
	for i, b := range out {
		if !isVowel(b) {
			continue
		}
		for !isVowel(v[j]) {
			j++
		}
		out[i] = v[j]
		j++
	}
	return string(out)
}
