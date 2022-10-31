package string

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestString(t *testing.T) {
	suites := []suite{
		{
			6,
			3,
		},
		{
			1,
			1,
		},
	}

	for _, s := range suites {
		ret := magicalString(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
