package of2

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret bool
}

func TestResults(t *testing.T) {
	suites := []suite{
		{
			1,
			true,
		},
		{
			10,
			false,
		},
		{
			16,
			true,
		},
		{
			24,
			false,
		},
		{
			46,
			true,
		},
	}

	for _, s := range suites {
		ret := reorderedPowerOf2(s.n)
		if s.ret != ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
