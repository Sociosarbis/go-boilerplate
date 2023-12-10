package stairs

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestStairs(t *testing.T) {
	suites := []suite{
		{
			2,
			2,
		},
		{
			3,
			3,
		},
	}

	for _, s := range suites {
		ret := climbStairs(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
