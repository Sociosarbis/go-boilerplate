package string

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestString(t *testing.T) {
	suites := []suite{
		{
			"ABC",
			10,
		},
		{
			"ABA",
			8,
		},
		{
			"LEETCODE",
			92,
		},
	}

	for _, s := range suites {
		ret := uniqueLetterString(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
