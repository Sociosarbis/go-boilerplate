package swap

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	num int
	ret int
}

func TestSwap(t *testing.T) {
	suites := []suite{
		{
			2736,
			7236,
		},
		{
			9973,
			9973,
		},
	}

	for _, s := range suites {
		ret := maximumSwap(s.num)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
