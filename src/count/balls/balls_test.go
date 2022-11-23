package balls

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	lowLimit  int
	highLimit int
	ret       int
}

func TestBalls(t *testing.T) {
	suites := []suite{
		{
			1,
			10,
			2,
		},
		{
			5,
			15,
			2,
		},
		{
			19,
			28,
			2,
		},
	}

	for _, s := range suites {
		ret := countBalls(s.lowLimit, s.highLimit)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
