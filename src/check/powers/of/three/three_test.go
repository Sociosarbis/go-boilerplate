package three

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret bool
}

func TestThree(t *testing.T) {
	suites := []suite{
		{
			12,
			true,
		},
		{
			91,
			true,
		},
		{
			21,
			false,
		},
	}

	for _, s := range suites {
		ret := checkPowersOfThree(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
