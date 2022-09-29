package string

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s1  string
	s2  string
	ret bool
}

func TestString(t *testing.T) {
	suites := []suite{
		{
			"waterbottle",
			"erbottlewat",
			true,
		},
		{
			"aa",
			"aba",
			false,
		},
	}

	for _, s := range suites {
		ret := isFlipedString(s.s1, s.s2)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
