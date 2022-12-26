package homogenous

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestHomogenouse(t *testing.T) {
	suites := []suite{
		{
			"abbcccaa",
			13,
		},
		{
			"xy",
			2,
		},
		{
			"zzzzz",
			15,
		},
	}
	for _, s := range suites {
		ret := countHomogenous(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
