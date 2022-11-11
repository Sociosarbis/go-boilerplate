package alikes

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret bool
}

func TestAlikes(t *testing.T) {
	suites := []suite{
		{
			"book",
			true,
		},
		{
			"textbook",
			false,
		},
	}

	for _, s := range suites {
		ret := halvesAreAlike(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
