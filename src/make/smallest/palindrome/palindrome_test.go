package palindrome

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret string
}

func TestPalindrome(t *testing.T) {
	suites := []suite{
		{
			"egcfe",
			"efcfe",
		},
		{
			"abcd",
			"abba",
		},
		{
			"seven",
			"neven",
		},
	}

	for _, s := range suites {
		ret := makeSmallestPalindrome(s.s)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
