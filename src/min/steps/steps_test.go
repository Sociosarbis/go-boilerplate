package steps

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestResults(t *testing.T) {
	suites := []suite{
		{
			3,
			3,
		},
		{
			1,
			0,
		},
	}

	for _, s := range suites {
		ret := minSteps(s.n)
		if s.ret != ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
