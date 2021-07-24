package time

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	time string
	ret  string
}

func TestMaximumTime(t *testing.T) {
	suites := []suite{
		{
			"2?:?0",
			"23:50",
		},
		{
			"0?:3?",
			"09:39",
		},
		{
			"1?:22",
			"19:22",
		},
	}

	for _, s := range suites {
		ret := maximumTime(s.time)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
