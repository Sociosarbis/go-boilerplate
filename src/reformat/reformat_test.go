package reformat

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret string
}

func TestReformat(t *testing.T) {
	suites := []suite{
		{
			"a0b1c2",
			"0a1b2c",
		},
		{
			"leetcode",
			"",
		},
		{
			"1229857369",
			"",
		},
	}

	for _, s := range suites {
		ret := reformat(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
