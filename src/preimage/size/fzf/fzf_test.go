package fzf

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	k   int
	ret int
}

func TestFzf(t *testing.T) {
	suites := []suite{
		{
			0,
			5,
		},
		{
			5,
			0,
		},
		{
			3,
			5,
		},
	}
	for _, s := range suites {
		ret := preimageSizeFZF(s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
