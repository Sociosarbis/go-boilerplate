package days

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestDays(t *testing.T) {
	suites := []suite{
		{
			10,
			4,
		},
		{
			6,
			3,
		},
	}

	for _, s := range suites {
		ret := minDays(s.n)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
