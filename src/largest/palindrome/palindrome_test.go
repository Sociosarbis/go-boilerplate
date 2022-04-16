package palindrome

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestPalindrome(t *testing.T) {
	suites := []suite{
		{
			2,
			987,
		},
		{
			1,
			9,
		},
	}

	for _, s := range suites {
		ret := largestPalindrome(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
