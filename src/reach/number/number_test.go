package number

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	target int
	ret    int
}

func TestNumber(t *testing.T) {
	suites := []suite{
		{
			2,
			3,
		},
		{
			3,
			2,
		},
	}

	for _, s := range suites {
		ret := reachNumber(s.target)
		if s.ret != ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
