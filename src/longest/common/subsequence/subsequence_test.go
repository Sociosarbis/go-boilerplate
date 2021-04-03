package subsequence

import (
	"testing"
)

type suite struct {
	text1 string
	text2 string
	ret   int
}

func TestLongestCommonSubsequence(t *testing.T) {
	suites := []suite{
		{
			"abcde",
			"ace",
			3,
		},
		{
			"abc",
			"abc",
			3,
		},
		{
			"abc",
			"def",
			0,
		},
	}

	for _, s := range suites {
		ret := longestCommonSubsequence(s.text1, s.text2)
		if ret != s.ret {
			t.Errorf("Failed, expected %v but get %v", s.ret, ret)
		}
	}
}
