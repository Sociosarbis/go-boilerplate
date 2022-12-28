package length

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestLength(t *testing.T) {
	suites := []suite{
		{
			"ca",
			2,
		},
		{
			"cabaabac",
			0,
		},
		{
			"aabccabba",
			3,
		},
	}

	for _, s := range suites {
		ret := minimumLength(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
