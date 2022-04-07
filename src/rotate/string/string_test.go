package string

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s    string
	goal string
	ret  bool
}

func TestString(t *testing.T) {
	suites := []suite{
		{
			"abcde",
			"cdeab",
			true,
		},
		{
			"abcde",
			"abced",
			false,
		},
	}

	for _, s := range suites {
		ret := rotateString(s.s, s.goal)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
