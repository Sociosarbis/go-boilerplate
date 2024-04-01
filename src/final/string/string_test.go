package string

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret string
}

func TestString(t *testing.T) {
	suites := []suite{
		{
			"string",
			"rtsng",
		},
		{
			"poiinter",
			"ponter",
		},
	}

	for _, s := range suites {
		ret := finalString(s.s)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
