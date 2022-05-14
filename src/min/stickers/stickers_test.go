package stickers

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	stickers []string
	target   string
	ret      int
}

func TestStickers(t *testing.T) {
	suites := []suite{
		{
			[]string{"with", "example", "science"},
			"thehat",
			3,
		},
		{
			[]string{"notice", "possible"},
			"basicbasic",
			-1,
		},
	}

	for _, s := range suites {
		ret := minStickers(s.stickers, s.target)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
