package k

import (
	"sort"
)

func isValid(s string, term []byte, k int) bool {
	n := len(s)
	var j int
	for i := 0; i < n && k != 0; i++ {
		if s[i] == term[j] {
			j++
			if j >= len(term) {
				k--
				j = 0
			}
		}
	}
	return k == 0
}

func permute(chars []byte, temp []byte, collected []bool, n int, cb func(option []byte) bool) bool {
	if len(temp) == n {
		return cb(temp)
	}
	for i, c := range chars {
		if !collected[i] {
			collected[i] = true
			if permute(chars, append(temp, c), collected, n, cb) {
				return true
			}
			collected[i] = false
		}
	}
	return false
}

func longestSubsequenceRepeatedK(s string, k int) string {
	counter := [26]int{}
	for _, c := range s {
		counter[c-'a']++
	}
	chars := make([]byte, 0, 7)
	for i := 0; i < 26; i++ {
		t := counter[i] / k
		if t > 0 {
			for j := 0; j < t; j++ {
				chars = append(chars, byte(i+97))
			}
		}
	}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] > chars[j]
	})
	var ret string
	temp := make([]byte, 0, 7)
	collected := make([]bool, len(chars))
	for i := len(chars); i > 0; i-- {
		if permute(chars, temp, collected, i, func(o []byte) bool {
			if isValid(s, o, k) {
				ret = string(o)
				return true
			}
			return false
		}) {
			return ret
		}
	}
	return ""
}
