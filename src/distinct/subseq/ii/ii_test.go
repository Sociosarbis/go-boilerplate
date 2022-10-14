package ii

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestII(t *testing.T) {
	suites := []suite{
		{
			"abc",
			7,
		},
		{
			"aba",
			6,
		},
		{
			"aaa",
			3,
		},
	}

	for _, s := range suites {
		ret := distinctSubseqII(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
