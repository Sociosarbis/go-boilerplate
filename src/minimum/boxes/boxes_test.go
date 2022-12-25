package boxes

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestBoxes(t *testing.T) {
	suites := []suite{
		{
			3,
			3,
		},
		{
			4,
			3,
		},
		{
			10,
			6,
		},
	}

	for _, s := range suites {
		ret := minimumBoxes(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
